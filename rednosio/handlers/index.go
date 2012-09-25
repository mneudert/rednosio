package handlers

import (
    "html/template"
    "net/http"
)

var (
    indexTemplates = template.Must(template.ParseFiles(
        "templates/index.html",
    ))
)

func Index(w http.ResponseWriter, r *http.Request) {
    indexTemplates.ExecuteTemplate(w, "index.html", nil)
}