package handlers

import (
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func DownloadThumb(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("thumbs/downloads/" + r.FormValue("id") + ".png")
	if nil != err {
		createThumb("downloads", r.FormValue("id"))

		f, err = os.Open("thumbs/downloads/" + r.FormValue("id") + ".png")
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
	f, err := os.Open("thumbs/uploads/" + r.FormValue("id") + ".png")
	if nil != err {
		createThumb("uploads", r.FormValue("id"))

		f, err = os.Open("thumbs/uploads/" + r.FormValue("id") + ".png")
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
	n, err := os.Open(folder + "/" + id + ".png")
	if nil != err {
		return
	}

	i, _, err := image.Decode(n)
	if nil != err {
		return
	}

	ires := resize.Resize(140, 0, i, resize.Lanczos3)

	nres, err := os.OpenFile("thumbs/"+folder+"/"+id+".png", os.O_RDWR|os.O_CREATE, 0666)
	if nil != err {
		return
	}

	png.Encode(nres, ires)
}
