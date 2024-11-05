package api

import (
	"fmt"
	"net/http"

	"likemind/internal/api/middlware"
	"likemind/website/page"

	"github.com/a-h/templ"
)

func BootstrapServer() error {
	profileHandler := middlware.Join(
		templ.Handler(page.ProfilePage()),
		middlware.Auth,
		middlware.Logging,
	)
	http.Handle("/profile", profileHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("website/static"))))

	fmt.Println("Listening on :8080")
	return http.ListenAndServe(":8080", nil)
}
