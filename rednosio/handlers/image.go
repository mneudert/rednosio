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

func rednose(m image.Image, x, y, size int) image.Image {
    f, err := os.Open("static/rednose.png")
    if nil != err { return m }

    n, err := png.Decode(f)
    if nil != err { return m }

    mrgba := rgba(m)
    nres := resize.Resize(uint(size), 0, n, resize.Lanczos3)

    nx, ny := nres.Bounds().Dx(), nres.Bounds().Dy()

    b := mrgba.Bounds()
    mn := image.NewRGBA(b)

    cp := image.Point{nx/2, ny/2}
    np := image.Point{nx/2 - x, ny/2 - y}

    draw.Draw(mn, mrgba.Bounds(), mrgba, image.ZP, draw.Src)
    draw.DrawMask(mn, mrgba.Bounds(), nres, np, &circle{cp, size/2}, np, draw.Over)

    return mn
}

func rgba(m image.Image) *image.RGBA {
    if r, ok := m.(*image.RGBA); ok {
        return r
    }

    b := m.Bounds()
    r := image.NewRGBA(b)

    draw.Draw(r, b, image.Transparent, image.ZP, draw.Src)
    draw.Draw(r, b, m, image.ZP, draw.Src)

    return r
}

type circle struct {
    p image.Point
    r int
}

func (c *circle) ColorModel() color.Model {
    return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
    return image.Rect(c.p.X - c.r,
                      c.p.Y - c.r,
                      c.p.X + c.r,
                      c.p.Y + c.r)
}

func (c *circle) At(x, y int) color.Color {
    xx, yy, rr := float64(x-c.p.X) + 0.5,
                  float64(y-c.p.Y) + 0.5,
                  float64(c.r)

    if xx*xx+yy*yy < rr*rr {
        return color.Alpha{255}
    }

    return color.Alpha{0}
}