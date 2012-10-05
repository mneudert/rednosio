package handlers

import (
    "html/template"
    "net/http"
)

var (
    browseTemplates = template.Must(template.ParseFiles(
        "templates/browse_downloads.html",
        "templates/browse_uploads.html",
        "templates/footer.html",
        "templates/header.html",
    ))
)

func BrowseDownloads(w http.ResponseWriter, r *http.Request) {
    page := new(Page)
    page.NavDownloads = true
    page.PageTitle = "Browse Downloads"

    browseTemplates.ExecuteTemplate(w, "browse_downloads.html", page)
}

func BrowseUploads(w http.ResponseWriter, r *http.Request) {
    page := new(Page)
    page.NavUploads = true
    page.PageTitle = "Browse Uploads"

    browseTemplates.ExecuteTemplate(w, "browse_uploads.html", page)
}