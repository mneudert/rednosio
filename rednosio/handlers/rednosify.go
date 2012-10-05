package handlers

import (
    "html/template"
    "net/http"
    "os"
)

var (
    editTemplates = template.Must(template.ParseFiles(
        "templates/footer.html",
        "templates/header.html",
        "templates/rednosify.html",
    ))
)

type editParams struct {
    ImgId string
}

func Rednosify(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    _, err := os.Stat("uploads/" + id + ".png")

    if nil != err {
        http.Redirect(w, r, "/", 302)
        return;
    }

    editTemplates.ExecuteTemplate(w, "rednosify.html", editParams{ImgId: id})
}