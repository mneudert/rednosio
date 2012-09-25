package handlers

import (
    "html/template"
    "net/http"
)

var (
    editTemplates = template.Must(template.ParseFiles(
        "templates/rednosify.html",
    ))
)

func Rednosify(w http.ResponseWriter, r *http.Request) {
    editTemplates.ExecuteTemplate(w, "rednosify.html", nil)
}