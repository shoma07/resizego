package main

import (
	"C"
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"math"

	"golang.org/x/image/draw"
)

//export resize64
func resize64(blob *C.char, limit int) *C.char {
	dec, _ := base64.StdEncoding.DecodeString(C.GoString(blob))

	src, _, _ := image.Decode(bytes.NewReader(dec))

	rct := src.Bounds()
	width := rct.Dx()
	height := rct.Dy()

	dst := &image.RGBA{}

	if height >= width {
		f := float64((width * limit))
		w := math.Round(f / float64(height))
		dst = image.NewRGBA(image.Rect(0, 0, int(w), limit))
	} else {
		f := float64((limit * height))
		h := math.Round(f / float64(width))
		dst = image.NewRGBA(image.Rect(0, 0, limit, int(h)))
	}

	draw.CatmullRom.Scale(dst, dst.Bounds(), src, rct, draw.Over, nil)

	w := new(bytes.Buffer)

	jpeg.Encode(w, dst, &jpeg.Options{Quality: 100})

	return C.CString(base64.StdEncoding.EncodeToString(w.Bytes()))
}

func main() {}
