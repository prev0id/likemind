package api

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"likemind/internal/api/middlware"
	"likemind/website/page"
)

func BootstrapServer() error {
	profileHandler := middlware.Join(
		templ.Handler(page.ProfilePage()),
		middlware.Recover,
		middlware.Logging,
		middlware.Auth,
	)

	http.Handle("/profile", profileHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("website/static"))))

	fmt.Println("Listening on :8080")
	return http.ListenAndServe(":8080", nil)
}
