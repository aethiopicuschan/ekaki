package ekaki

import (
	"image"
	"io"

	"github.com/chai2010/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// ImageFrom reads an image from a reader
func ImageFrom(img io.Reader) (i image.Image, t Target, err error) {
	// gif, jpeg, png
	i, s, err := image.Decode(img)
	if err == nil {
		switch s {
		case "jpeg":
			t = TargetJPG
		case "png":
			t = TargetPNG
		case "gif":
			t = TargetGIF
		}
		return
	}

	// bmp
	i, err = bmp.Decode(img)
	if err == nil {
		t = TargetBMP
		return
	}

	// webp
	i, err = webp.Decode(img)
	if err == nil {
		t = TargetWebp
		return
	}

	// tiff
	i, err = tiff.Decode(img)
	if err == nil {
		t = TargetTiff
		return
	}

	err = ErrUnsupported

	return
}
