package ekaki

import (
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/gen2brain/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// Convert transforms the source image into the target format
func Convert(source io.Reader, target Target) (output []byte, err error) {
	// Convert the source image to an image.Image
	img, _, err := ImageFrom(source)
	if err != nil {
		return
	}

	// Convert the image.Image to the target format
	var buf bytes.Buffer
	switch target {
	case TargetBMP:
		err = bmp.Encode(&buf, img)
	case TargetJPG:
		err = jpeg.Encode(&buf, img, nil)
	case TargetPNG:
		err = png.Encode(&buf, img)
	case TargetGIF:
		err = gif.Encode(&buf, img, nil)
	case TargetWebp:
		err = webp.Encode(&buf, img)
	case TargetTiff:
		err = tiff.Encode(&buf, img, nil)
	default:
		err = ErrUnsupported
	}
	if err != nil {
		return
	}

	output = buf.Bytes()

	return
}
