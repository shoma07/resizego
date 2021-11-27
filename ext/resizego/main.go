package main

import (
	"C"
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"io"
	"math"

	"github.com/rwcarlsen/goexif/exif"
	"golang.org/x/image/draw"
)

//export resize64
func resize64(blob *C.char, limit int, quality int) *C.char {
	dec, _ := base64.StdEncoding.DecodeString(C.GoString(blob))

	origin, _, _ := image.Decode(bytes.NewReader(dec))
	orientation := readOrientation(bytes.NewReader(dec))
	src := applyOrientation(origin, orientation)

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

	jpeg.Encode(w, dst, &jpeg.Options{Quality: quality})

	return C.CString(base64.StdEncoding.EncodeToString(w.Bytes()))
}

func readOrientation(r io.Reader) int {
	e, err := exif.Decode(r)
	if err != nil {
		return 1
	}
	tag, err := e.Get(exif.Orientation)
	if err != nil {
		return 1
	}
	o, err := tag.Int(0)
	if err != nil {
		return 1
	}
	return o
}

func applyOrientation(i image.Image, o int) image.Image {
	switch o {
	case 2:
		return fixTopRight(i)
	case 3:
		return fixBottomRight(i)
	case 4:
		return fixBottomLeft(i)
	case 5:
		return fixLeftTop(i)
	case 6:
		return fixRightTop(i)
	case 7:
		return fixRightBottom(i)
	case 8:
		return fixLeftBottom(i)
	default:
		return i
	}
}

func fixTopRight(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(s)

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(x, y, i.At(s.Max.X-x-1, y))
		}
	}

	return o
}

func fixBottomRight(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(s)

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(x, y, i.At(s.Max.X-x-1, s.Max.Y-y-1))
		}
	}

	return o
}

func fixBottomLeft(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(s)

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(x, y, i.At(x, s.Max.Y-y-1))
		}
	}

	return o
}

func fixLeftTop(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(image.Rect(s.Min.Y, s.Min.X, s.Max.Y, s.Max.X))

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(y, x, i.At(x, y))
		}
	}

	return o
}

func fixRightTop(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(image.Rect(s.Min.Y, s.Min.X, s.Max.Y, s.Max.X))

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(y, x, i.At(x, s.Max.Y-y-1))
		}
	}

	return o
}

func fixRightBottom(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(image.Rect(s.Min.Y, s.Min.X, s.Max.Y, s.Max.X))

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(y, x, i.At(s.Max.X-x-1, s.Max.Y-y-1))
		}
	}

	return o
}

func fixLeftBottom(i image.Image) image.Image {
	s := i.Bounds()
	o := image.NewRGBA(image.Rect(s.Min.Y, s.Min.X, s.Max.Y, s.Max.X))

	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			o.Set(y, x, i.At(s.Max.X-x-1, y))
		}
	}

	return o
}

func main() {}
