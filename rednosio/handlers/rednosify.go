package handlers

import (
	"html/template"
	"net/http"
	"os"
)

var (
	rednosifyTemplates = template.Must(template.ParseFiles(
		"templates/footer.html",
		"templates/header.html",
		"templates/rednosify.html",
	))
)

func Rednosify(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, err := os.Stat("uploads/" + id + ".png")

	if nil != err {
		http.Redirect(w, r, "/", 302)
		return
	}

	page := new(RednosifyPage)
	page.ImgId = id
	page.NavHome = true
	page.PageTitle = "Rednosify Image"

	rednosifyTemplates.ExecuteTemplate(w, "rednosify.html", page)
}
