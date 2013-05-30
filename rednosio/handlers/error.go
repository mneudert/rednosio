package handlers

import (
	"html/template"
	"net/http"
)

var (
	errorTemplates = template.Must(template.ParseFiles(
		"templates/error.html",
	))
)

func Error(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				w.WriteHeader(http.StatusInternalServerError)
				errorTemplates.ExecuteTemplate(w, "error.html", err)
			}
		}()

		fn(w, r)
	}
}
