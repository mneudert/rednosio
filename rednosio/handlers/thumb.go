package handlers

import (
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func DownloadThumb(w http.ResponseWriter, r *http.Request) {
	thumb  := fmt.Sprintf("thumbs/downloads/%s.png", r.FormValue("id"))
	f, err := os.Open(thumb)
	if nil != err {
		createThumb("downloads", r.FormValue("id"))

		f, err = os.Open(thumb)
		if nil != err {
			return
		}
	}

	i, _, err := image.Decode(f)
	if nil != err {
		return
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, i)
}

func UploadThumb(w http.ResponseWriter, r *http.Request) {
	thumb  := fmt.Sprintf("thumbs/uploads/%s.png", r.FormValue("id"))
	f, err := os.Open(thumb)
	if nil != err {
		createThumb("uploads", r.FormValue("id"))

		f, err = os.Open(thumb)
		if nil != err {
			return
		}
	}

	i, _, err := image.Decode(f)
	if nil != err {
		return
	}

	w.Header().Set("Content-type", "image/png")
	png.Encode(w, i)
}

func createThumb(folder, id string) {
	thumb  := fmt.Sprintf("%s/%s.png", folder, id)
	n, err := os.Open(thumb)
	if nil != err {
		return
	}

	i, _, err := image.Decode(n)
	if nil != err {
		return
	}

	ires := resize.Resize(140, 0, i, resize.Lanczos3)

	nres, err := os.OpenFile("thumbs/" + thumb, os.O_RDWR|os.O_CREATE, 0666)
	if nil != err {
		return
	}

	png.Encode(nres, ires)
}
