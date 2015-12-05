package processor

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"math"

	"code.google.com/p/graphics-go/graphics"
	"code.google.com/p/graphics-go/graphics/interp"
	"github.com/rwcarlsen/goexif/exif"
)

var affines map[int]graphics.Affine = map[int]graphics.Affine{
	1: graphics.I,
	2: graphics.I.Scale(-1, 1),
	3: graphics.I.Scale(-1, -1),
	4: graphics.I.Scale(1, -1),
	5: graphics.I.Rotate(toRadian(90)).Scale(-1, 1),
	6: graphics.I.Rotate(toRadian(90)),
	7: graphics.I.Rotate(toRadian(-90)).Scale(-1, 1),
	8: graphics.I.Rotate(toRadian(-90)),
}

// Process returns an image that is applied Exif orientation tag.
func Process(r io.Reader) (d image.Image, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	s, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return
	}
	o, err := ReadOrientation(bytes.NewReader(b))
	if err != nil {
		// When orientation can't be read, return original image.
		return s, nil
	}
	d = ApplyOrientation(s, o)
	return
}

// ReadOrientation returns Exif orientation tag.
func ReadOrientation(r io.Reader) (o int, err error) {
	e, err := exif.Decode(r)
	if err != nil {
		return
	}
	tag, err := e.Get(exif.Orientation)
	if err != nil {
		return
	}
	o, err = tag.Int(0)
	if err != nil {
		return
	}
	return
}

// ApplyOrientation applies orientation to image.
func ApplyOrientation(s image.Image, o int) (d draw.Image) {
	bounds := s.Bounds()
	// Swap width and height when orientation between 5 and 8
	if o >= 5 && o <= 8 {
		bounds = rotateRect(bounds)
	}
	d = image.NewRGBA64(bounds)
	affine := affines[o]
	affine.TransformCenter(d, s, interp.Bilinear)
	return
}

func toRadian(d float64) float64 {
	return math.Pi * d / 180
}

func rotateRect(r image.Rectangle) image.Rectangle {
	s := r.Size()
	return image.Rectangle{r.Min, image.Point{s.Y, s.X}}
}
