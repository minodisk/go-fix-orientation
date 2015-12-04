package processor_test

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"testing"

	"github.com/minodisk/go-fix-orientation/processor"
)

func TestProcess(t *testing.T) {
	b, err := ioutil.ReadFile("./fixtures/f.jpg")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := jpeg.Decode(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range []string{
		"./fixtures/f-orientation-1.jpg",
		"./fixtures/f-orientation-2.jpg",
		"./fixtures/f-orientation-3.jpg",
		"./fixtures/f-orientation-4.jpg",
		"./fixtures/f-orientation-5.jpg",
		"./fixtures/f-orientation-6.jpg",
		"./fixtures/f-orientation-7.jpg",
		"./fixtures/f-orientation-8.jpg",
		"./fixtures/f-png8.png",
		"./fixtures/f-png24.png",
		"./fixtures/f.gif",
		"./fixtures/f.jpg",
	} {
		f, err := os.Open(p)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := processor.Process(f)
		if err != nil {
			t.Fatal(err)
		}
		if !nearPixels(actual, expected) {
			t.Fail()
		}
	}
}

func nearPixels(a, b image.Image) bool {
	bounds := a.Bounds()
	if !bounds.Eq(b.Bounds()) {
		return false
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ar, ag, ab, aa := a.At(x, y).RGBA()
			br, bg, bb, ba := b.At(x, y).RGBA()
			if !nearPixel(ar, ag, ab, aa, br, bg, bb, ba) {
				return false
			}
		}
	}
	return true
}

func nearPixel(ar, ag, ab, aa, br, bg, bb, ba uint32) bool {
	return near(ar, br) && near(ag, bg) && near(ab, bb) && near(aa, ba)
}

func near(m, n uint32) bool {
	return diff(m, n) < math.MaxUint16
}

func diff(m, n uint32) uint32 {
	if m > n {
		return m - n
	}
	return n - m
}
