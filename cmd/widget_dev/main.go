package main

import (
	"log"
	"net/http"

	"likemind/internal/api/handlers/widget_dev"
)

func main() {
	w := widget_dev.New()
	http.HandleFunc("/", w.HandleTestingPage)

	http.HandleFunc("/dev/tests", w.HandleListTests)
	http.HandleFunc("/dev/test/{widget}", w.HandleWidget)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("website/static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
