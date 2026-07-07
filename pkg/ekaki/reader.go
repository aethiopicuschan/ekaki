package ekaki

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/gen2brain/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// ImageFrom reads an image from a reader.
func ImageFrom(img io.Reader) (i image.Image, t Target, err error) {
	data, err := io.ReadAll(img)
	if err != nil {
		return nil, t, err
	}

	// webp
	//
	// DecodeAll works for both still WebP and animated WebP.
	// For animated WebP, return only the first frame.
	i, err = decodeFirstWebPFrame(data)
	if err == nil {
		return i, TargetWebp, nil
	}

	// gif, jpeg, png
	i, s, err := image.Decode(bytes.NewReader(data))
	if err == nil {
		switch s {
		case "jpeg":
			t = TargetJPG
		case "png":
			t = TargetPNG
		case "gif":
			t = TargetGIF
		case "bmp":
			t = TargetBMP
		case "tiff":
			t = TargetTiff
		case "webp":
			t = TargetWebp
		default:
			err = ErrUnsupported
			return
		}
		return
	}

	// bmp
	i, err = bmp.Decode(bytes.NewReader(data))
	if err == nil {
		t = TargetBMP
		return
	}

	// tiff
	i, err = tiff.Decode(bytes.NewReader(data))
	if err == nil {
		t = TargetTiff
		return
	}

	err = ErrUnsupported
	return
}

func decodeFirstWebPFrame(data []byte) (image.Image, error) {
	anim, err := webp.DecodeAll(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if anim == nil || len(anim.Image) == 0 || anim.Image[0] == nil {
		return nil, errors.New("webp: no frame found")
	}

	return anim.Image[0], nil
}
