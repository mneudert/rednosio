package handlers

import (
    "image"
    "image/png"
    "net/http"
    "os"
)

func Image(w http.ResponseWriter, r *http.Request) {
    f, err := os.Open("uploads/" + r.FormValue("id") + ".png")
    if nil != err { return; }

    i, _, err := image.Decode(f)
    if nil != err { return; }

    w.Header().Set("Content-type", "image/png")
    png.Encode(w, i)
}