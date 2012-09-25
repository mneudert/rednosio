package handlers

import (
    "errors"
    "html/template"
    "net/http"
)

var (
    indexTemplates = template.Must(template.ParseFiles(
        "templates/index.html",
    ))
)

type indexParams struct {
    ErrMsg string
}

func Index(w http.ResponseWriter, r *http.Request) {
    if "POST" == r.Method {
        err := handleUpload(w, r)

        if nil == err {
            http.Redirect(w, r, "/rednosify", 302)
            return;
        }

        indexTemplates.ExecuteTemplate(w, "index.html", indexParams{ErrMsg: err.Error()})
        return;
    }

    indexTemplates.ExecuteTemplate(w, "index.html", nil)
}

func handleUpload(w http.ResponseWriter, r *http.Request) error {
    _, _, err := r.FormFile("image")

    if nil != err {
        return errors.New("No file selected!")
    }

    return nil
}