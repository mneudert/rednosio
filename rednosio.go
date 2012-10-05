package main

import (
    "net/http"
    "./rednosio/handlers"
)

func main() {
    http.HandleFunc("/", handlers.Error(handlers.Index))
    http.HandleFunc("/browse/downloads", handlers.Error(handlers.BrowseDownloads));
    http.HandleFunc("/browse/uploads", handlers.Error(handlers.BrowseUploads));
    http.HandleFunc("/image", handlers.Error(handlers.Image))
    http.HandleFunc("/rednosify", handlers.Error(handlers.Rednosify))
    http.HandleFunc("/save", handlers.Error(handlers.SaveImage))

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    http.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir("./downloads"))))

    http.ListenAndServe(":8080", nil)
}