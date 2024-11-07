package main

import (
	"likemind/internal/api/handlers/widget_dev"
	"log"
	"net/http"
)

func main() {
	w := widget_dev.New()
	http.HandleFunc("/", w.HandleTestingPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("website/static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
