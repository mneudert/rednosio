package handlers

import (
    "crypto/sha1"
    "errors"
    "fmt"
    "html/template"
    "image"
    "image/png"
    _ "image/jpeg"
    "io"
    "io/ioutil"
    "net/http"
    "os"
)

var (
    indexTemplates = template.Must(template.ParseFiles(
        "templates/footer.html",
        "templates/header.html",
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
            return;
        }

        indexTemplates.ExecuteTemplate(w, "index.html", indexParams{ErrMsg: err.Error()})
        return;
    }

    indexTemplates.ExecuteTemplate(w, "index.html", nil)
}

func handleUpload(w http.ResponseWriter, r *http.Request) error {
    f, _, err := r.FormFile("image")
    if nil != err { return errors.New("No file selected.") }

    t, err := ioutil.TempFile("uploads", "temp-")
    if nil != err { return errors.New("Could not create temp file.") }

    _, err = io.Copy(t, f)
    if nil != err { return errors.New("Could not copy upload to temp file.") }

    fh, err := os.Open(t.Name())
    if nil != err { return errors.New("Could not open temp file.") }

    fc, err := os.Open(t.Name())
    if nil != err { return errors.New("Could not open temp file.") }

    mc, _, err := image.Decode(fc)
    if nil != err {
        os.Remove(t.Name())
        return errors.New("Could not decode temp file.")
    }

    h := sha1.New()
    _, err = io.Copy(h, fh)
    sha1 := fmt.Sprintf("%x", h.Sum(nil))

    _, err = os.Stat("uploads/" + sha1 + ".png")

    if nil != err {
        fp, err := os.OpenFile("uploads/" + sha1 + ".png", os.O_RDWR | os.O_CREATE, 0666)
        if nil != err { return errors.New("Could not create upload file.") }

        encerr := png.Encode(fp, mc)
        if nil != encerr {
            os.Remove(t.Name())
            os.Remove(fp.Name())
            return errors.New("Could not convert temp file to PNG.")
        }
    }

    os.Remove(t.Name())

    http.Redirect(w, r, "/rednosify?id=" + sha1, 302)

    return nil
}