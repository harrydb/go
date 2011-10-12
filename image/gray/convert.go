package gray

import (
	"image"
	"math"
)

type ConvertFunc func (image.Color) image.GrayColor

func Convert(m image.Image, convertColor ConvertFunc) *image.Gray {
	b := m.Bounds()
	gray := image.NewGray(b.Dx(), b.Dy())
	pos := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			gray.Pix[pos] = convertColor(m.At(x, y)).Y
			pos++
		}
	}
	return gray
}

// ToGrayAverage converts image.Color c to grayscale using the average of the RGB components.
// 
// The formula used for conversion is: Y = (r + g + b) / 3.
func ToGrayAverage(c image.Color) image.GrayColor {
	r, g, b, _ := c.RGBA()
	Y := (10*(r + g + b) + 5) / 10
	return image.GrayColor{uint8(Y >> 8)}
}

// ToGrayLuma converts image.Color c to grayscale using NTSC weighting.
//
// The formula used for conversion is: Y = 0.299*r + 0.587*g + 0.114*b.
// The same formula is used by image.GrayColorModel.Convert().
func ToGrayLuma(c image.Color) image.GrayColor {
	r, g, b, _ := c.RGBA()
	Y := (299*r + 587*g + 114*b + 500) / 1000
	return image.GrayColor{uint8(Y >> 8)}
}

// ToGrayLuma709 converts image.Color c to grayscale using Rec 709 weighting.
//
// The formula used for conversion is: Y = 0.299*r + 0.587*g + 0.114*b.
// The same formula is used by image.GrayColorModel.Convert().
func ToGrayLuma709(c image.Color) image.GrayColor {
	r, g, b, _ := c.RGBA()
	Y := (2125*r + 7154*g + 721*b + 5000) / 10000
	return image.GrayColor{uint8(Y >> 8)}
}

// ToGrayLuminance converts image.Color c to grayscale using Rec 709.
//
// The formula used for conversion is: Y' = 0.2125*R' + 0.7154*G' + 0.0721*B'
// where r, g and b are gamma expanded with gamma 2.2 and final Y is Y'
// gamma compressed again.
// The same formula is used by image.GrayColorModel.Convert().
func ToGrayLuminance(c image.Color) image.GrayColor {
	rr, gg, bb, _ := c.RGBA()
	r := math.Pow(float64(rr), 2.2)
	g := math.Pow(float64(gg), 2.2)
	b := math.Pow(float64(bb), 2.2)
	y := math.Pow(0.2125*r + 0.7154*g + 0.0721*b, 1/2.2)
	Y := uint16(y + 0.5)
	return image.GrayColor{uint8(Y >> 8)}
}

// ToGrayLightness converts image.Color c to grayscale using the lightness.
//
// The formula used for conversion is: Y = (max(r,g,b) + min(r,g,b)) / 2.
func ToGrayLightness(c image.Color) image.GrayColor {
	r, g, b, _ := c.RGBA()
	max := Max(r, Max(g, b))
	min := Min(r, Min(g, b))
	Y := (10*(min + max) + 5) / 20
	return image.GrayColor{uint8(Y >> 8)}
}

// ToGrayHSV converts image.Color c to grayscale using the V component in HSV color space.
//
// The formula used for conversion is: Y = max(r,g,b).
func ToGrayValue(c image.Color) image.GrayColor {
	r, g, b, _ := c.RGBA()
	Y := Max(r, Max(g, b))
	return image.GrayColor{uint8(Y >> 8)}
}

// ToGrayRed converts image.Color c to grayscale using the R component.
func ToGrayRed(c image.Color) image.GrayColor {
	r, _, _, _ := c.RGBA()
	return image.GrayColor{uint8(r >> 8)}
}

// ToGrayGreen converts image.Color c to grayscale using the G component.
func ToGrayGreen(c image.Color) image.GrayColor {
	_, g, _, _ := c.RGBA()
	return image.GrayColor{uint8(g >> 8)}
}

// ToGrayBlue converts image.Color c to grayscale using the B component.
func ToGrayBlue(c image.Color) image.GrayColor {
	_, _, b, _ := c.RGBA()
	return image.GrayColor{uint8(b >> 8)}
}

// ToGrayAlpha converts image.Color c to grayscale using the A component.
func ToGrayAlpha(c image.Color) image.GrayColor {
	_, _, _, a := c.RGBA()
	return image.GrayColor{uint8(a >> 8)}
}

func Max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
