package api_handler

import (
	"net/http"

	"likemind/internal/app/middleware"
	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	auth        middleware.Middleware
	authService auth.Service
	profile     profile.Service
}

func New(profileSvc profile.Service, authService auth.Service, authMiddleware middleware.Middleware) *ApiHandler {
	return &ApiHandler{
		auth:        authMiddleware,
		profile:     profileSvc,
		authService: authService,
	}
}

func (h *ApiHandler) Prefix() string {
	return domain.PathPrefixAPI
}

func (h *ApiHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Group(func(public chi.Router) {
		public.Post(domain.PathSignUp, h.SignUp)
		public.Post(domain.PathSignIn, h.SignIn)
	})

	router.Group(func(protected chi.Router) {
		protected.Use(h.auth)

		protected.Post(domain.PathLogOut, h.LogOut)
	})

	return router
}

func (h *ApiHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if userID, err := h.authService.ValidateSessionCookie(w, r); err == nil {
		h.redirectToProfile(userID, w, r)
		return
	}

	request, err := httpin.Decode[SignInRequst](r)
	if err != nil {
		common.ServeError(w, r, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	creds, err := h.authService.Signin(
		ctx,
		domain.Email(request.Email),
		domain.Password(request.Password),
	)
	if err != nil {
		common.ServeError(w, r, err, http.StatusBadRequest)
		return
	}

	if err := h.authService.SetSessionCookie(creds.ID, w, r); err != nil {
		common.ServeError(w, r, err, http.StatusInternalServerError)
		return
	}

	h.redirectToProfile(creds.UserID, w, r)
}

func (h *ApiHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if userID, err := h.authService.ValidateSessionCookie(w, r); err == nil {
		h.redirectToProfile(userID, w, r)
		return
	}

	request, err := httpin.Decode[SignUpRequest](r)
	if err != nil {
		common.ServeError(w, r, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	userID, err := h.profile.CreateUser(ctx, request.getCreateUserRequest())
	if err != nil {
		common.ServeError(w, r, err, http.StatusBadRequest)
		return
	}

	uuid, err := h.authService.NewUserCredentials(
		ctx,
		userID,
		domain.Email(request.Email),
		domain.Password(request.Password),
	)
	if err != nil {
		common.ServeError(w, r, err, http.StatusInternalServerError)
		return
	}

	if err := h.authService.SetSessionCookie(uuid, w, r); err != nil {
		common.ServeError(w, r, err, http.StatusInternalServerError)
		return
	}

	h.redirectToProfile(userID, w, r)
}

func (h *ApiHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	h.authService.InvalidateSessionCookie(w, r)
	http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
}

func (h *ApiHandler) redirectToProfile(userID int64, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile, err := h.profile.GetProfile(ctx, userID)
	if err != nil {
		// corner case
		searchUsersPath := common.SetPathVariables(domain.PathSearch, common.PathVars{domain.PathVarType: domain.TypeUser})
		common.Redirect(w, searchUsersPath)
		return
	}

	profilePath := common.SetPathVariables(domain.PathUserPage, common.PathVars{domain.PathVarNickname: profile.User.Nickname})
	common.Redirect(w, profilePath)
}
