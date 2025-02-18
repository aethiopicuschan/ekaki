package ekaki

import (
	"strings"
)

// Target represents the target format of an image
type Target string

const (
	TargetBMP  Target = "bmp"
	TargetJPG  Target = "jpg"
	TargetPNG  Target = "png"
	TargetGIF  Target = "gif"
	TargetWebp Target = "webp"
	TargetTiff Target = "tiff"
)

// SupportedTargets returns a list of supported target formats
func SupportedTargets() []Target {
	return []Target{
		TargetBMP,
		TargetJPG,
		TargetPNG,
		TargetGIF,
		TargetWebp,
		TargetTiff,
	}
}

// TargetFromExpr returns the target format from a given string
func TargetFromExpr(expr string) (t Target, err error) {
	expr = strings.TrimPrefix(expr, ".")
	switch expr {
	case "bmp":
		t = TargetBMP
	case "jpg", "jpeg":
		t = TargetJPG
	case "png":
		t = TargetPNG
	case "gif":
		t = TargetGIF
	case "webp":
		t = TargetWebp
	case "tiff":
		t = TargetTiff
	default:
		err = ErrUnsupported
	}
	return
}
