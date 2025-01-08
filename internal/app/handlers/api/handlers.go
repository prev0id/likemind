package api_handler

import (
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	authSvc auth.Service
	profile profile.Service
}

func New(authSvc auth.Service, profileSvc profile.Service) *ApiHandler {
	return &ApiHandler{
		authSvc: authSvc,
		profile: profileSvc,
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
		protected.Use(h.authSvc.Middleware)

		protected.Post(domain.PathLogOut, h.LogOut)
	})

	return router
}

func (h *ApiHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if userID, err := h.authSvc.ValidateSessionCookie(w, r); err == nil {
		h.redirectToProfile(userID, w, r)
		return
	}

	request, err := httpin.Decode[SignInRequst](r)
	if err != nil {
		common.ServeError(w, r, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	creds, err := h.authSvc.ValidateCredentials(ctx, request.Email, request.Password)
	if err != nil {
		common.ServeError(w, r, err, http.StatusBadRequest)
		return
	}

	if err := h.authSvc.SetSessionCookie(creds.UUID, w, r); err != nil {
		common.ServeError(w, r, err, http.StatusInternalServerError)
		return
	}

	h.redirectToProfile(creds.UserID, w, r)
}

func (h *ApiHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if userID, err := h.authSvc.ValidateSessionCookie(w, r); err == nil {
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

	uuid, err := h.authSvc.SetCredentials(ctx, userID, request.Email, request.Password)
	if err != nil {
		common.ServeError(w, r, err, http.StatusInternalServerError)
		return
	}

	if err := h.authSvc.SetSessionCookie(uuid, w, r); err != nil {
		common.ServeError(w, r, err, http.StatusInternalServerError)
		return
	}

	h.redirectToProfile(userID, w, r)
}

func (h *ApiHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	h.authSvc.InvalidateSessionCookie(w, r)
	http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
}

func (h *ApiHandler) redirectToProfile(userID int64, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile, err := h.profile.GetUser(ctx, userID)
	if err != nil {
		// corner case
		searchUsersPath := common.SetPathVariables(domain.PathSearch, common.PathVars{domain.PathVarType: domain.TypeUser})
		common.Redirect(w, searchUsersPath)
		return
	}

	profilePath := common.SetPathVariables(domain.PathUserPage, common.PathVars{domain.PathVarNickname: profile.Nickname})
	common.Redirect(w, profilePath)
}
