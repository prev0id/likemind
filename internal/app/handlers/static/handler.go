package static_handler

import (
	"log"
	"net/http"

	"likemind/internal/domain"
	"likemind/website/static"

	"github.com/go-chi/chi/v5"
)

type StaticHandler struct{}

func New() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) Prefix() string {
	return domain.PathPrefixStatic
}

func (h *StaticHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get(domain.PatternFile, h.GetFile)

	return router
}

func (h *StaticHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	descriptor, ok := static.Files[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", descriptor.MimeType)
	if _, err := w.Write(descriptor.File); err != nil {
		log.Printf("err writing static file %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
