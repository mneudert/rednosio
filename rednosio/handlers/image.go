package handlers

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "net/http"
    "os"
    "strconv"

    "github.com/nfnt/resize"
)

func Image(w http.ResponseWriter, r *http.Request) {
    f, err := os.Open("uploads/" + r.FormValue("id") + ".png")
    if nil != err { return; }

    i, _, err := image.Decode(f)
    if nil != err { return; }

    get := func(n string) int {
        i, _ := strconv.Atoi(r.FormValue(n))
        return i
    }

    x, y, s := get("x"), get("y"), get("s")

    if x > 0 && y > 0 && s > 0 {
        i = rednose(i, x, y, s)
    }

    w.Header().Set("Content-type", "image/png")
    png.Encode(w, i)
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
    sha1 := r.FormValue("id")

    f, err := os.Open("uploads/" + sha1 + ".png")
    if nil != err {
        http.Redirect(w, r, "/", 302)
        return
    }

    i, _, err := image.Decode(f)
    if nil != err {
        http.Redirect(w, r, "/rednosify?id=" + sha1, 302)
        return
    }

    get := func(n string) int {
        i, _ := strconv.Atoi(r.FormValue(n))
        return i
    }

    x, y, s := get("x"), get("y"), get("s")

    if x > 0 && y > 0 && s > 0 {
        i = rednose(i, x, y, s)
    }

    d, err := os.OpenFile("downloads/" + r.FormValue("id") + "_" +
                                         r.FormValue("x") + "_" +
                                         r.FormValue("y") + "_" +
                                         r.FormValue("s") + ".png",
                          os.O_RDWR | os.O_CREATE, 0666)
    if nil != err {
        http.Redirect(w, r, "/rednosify?id=" + sha1, 302)
        return
    }

    png.Encode(d, i)

    http.Redirect(w, r, "downloads/" + r.FormValue("id") + "_" +
                                       r.FormValue("x") + "_" +
                                       r.FormValue("y") + "_" +
                                       r.FormValue("s") + ".png", 302)
}

func rednose(m image.Image, x, y, size int) image.Image {
    n := getRednose(size)
    if n == nil { return m }

    b := m.Bounds()
    mn := image.NewRGBA(b)

    cp := image.Point{size/2, size/2}
    np := image.Point{size/2 - x, size/2 - y}

    draw.Draw(mn, m.Bounds(), m, image.ZP, draw.Src)
    draw.DrawMask(mn, m.Bounds(), n, np, &Circle{cp, size/2}, np, draw.Over)

    return mn
}

type Circle struct {
    p image.Point
    r int
}

func (c *Circle) ColorModel() color.Model {
    return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
    return image.Rect(c.p.X - c.r,
                      c.p.Y - c.r,
                      c.p.X + c.r,
                      c.p.Y + c.r)
}

func (c *Circle) At(x, y int) color.Color {
    xx, yy, rr := float64(x-c.p.X) + 0.5,
                  float64(y-c.p.Y) + 0.5,
                  float64(c.r)

    if xx*xx+yy*yy < rr*rr {
        return color.Alpha{255}
    }

    return color.Alpha{0}
}

func getRednose(size int) image.Image {
    s := strconv.Itoa(size)

    _, err := os.Stat("static/rednose/" + s + ".png")

    if nil != err {
        n, err := os.Open("static/rednose.png")
        if nil != err { return nil }

        i, _, err := image.Decode(n)
        if nil != err { return nil }

        ires := resize.Resize(uint(size), 0, i, resize.Lanczos3)

        nres, err := os.OpenFile("static/rednose/" + s + ".png", os.O_RDWR | os.O_CREATE, 0666)
        if nil != err { return nil }

        png.Encode(nres, ires)
    }

    fs, err := os.Open("static/rednose/" + s + ".png")
    if nil != err { return nil }

    ns, err := png.Decode(fs)
    if nil != err { return nil }

    return ns
}