package main

import (
    "net/http"
    "./rednosio/handlers"
)

func main() {
    http.HandleFunc("/", handlers.Error(handlers.Index))
    http.HandleFunc("/rednosify", handlers.Error(handlers.Rednosify))
    http.HandleFunc("/image", handlers.Error(handlers.Image))

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    http.ListenAndServe(":8080", nil)
}