package transform

import (
	"math"
	"image"
	"image/color"
	"fmt"
)

type TransformFunc func(int, int) (float64, float64)
type InterpolationFunc func(image.Image, TransformFunc, int, int, int) image.Image

const (
	BORDER_TRANSPARENT = 0
	BORDER_COPY        = 1
)

type WritableImage interface {
	image.Image
	Set(x, y int, c color.Color)
}

func Nearest(src image.Image, transform TransformFunc, width, height, borderMethod int) image.Image {
	dst := newImage(src, width, height)
	b := src.Bounds()
	for ydst := 0; ydst < height; ydst++ {
		for xdst := 0; xdst < width; xdst++ {
			X, Y := transform(xdst, ydst)
			x := int(X) + b.Min.X
			y := int(Y) + b.Min.Y
			if (X) < 0.0 {
				x -= 1
			}
			if (Y) < 0.0 {
				y -= 1
			}
			// Is this pixel outside the source image?
			if x < b.Min.X || y < b.Min.Y || x >= b.Max.X || y >= b.Max.Y {
				continue
			}
			dst.Set(xdst, ydst, getColor(src, x, y, borderMethod))
		}
	}
	return dst
}

func Bilinear(src image.Image, transform TransformFunc, width, height, borderMethod int) image.Image {
	// Fast path for certain image types.
	switch img := src.(type) {
	case *image.RGBA:
		return bilinearRGBA(img, transform, width, height, borderMethod)
	case *image.Gray:
		return bilinearGray(img, transform, width, height, borderMethod)
	}
	// Standard path
	dst := newImage(src, width, height)
	b := src.Bounds()

	for ydst := 0; ydst < height; ydst++ {
		for xdst := 0; xdst < width; xdst++ {
			X, Y := transform(xdst, ydst)
			X -= 0.5
			Y -= 0.5
			x := int(X)
			y := int(Y)
			if X < 0.0 {
				x -= 1
			}
			if Y < 0.0 {
				y -= 1
			}
			// Are all neighbours outside the source image?
			if x < -1 || y < -1 || x >= b.Dx() || y >= b.Dy() {
				continue
			}
			// Pixel weights
			dx := X - float64(x)
			dy := Y - float64(y)
			// Image boundaries
			x += b.Min.X
			y += b.Min.Y
			// Calculate new color, add 0.5 to round to nearest integer.
			var R, G, B, A float64 = 0.5, 0.5, 0.5, 0.5
			r, g, b, a := getColor(src, x, y, borderMethod).RGBA()
			w := (1 - dx) * (1 - dy)
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			r, g, b, a = getColor(src, x+1, y, borderMethod).RGBA()
			w = dx * (1 - dy)
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			r, g, b, a = getColor(src, x, y+1, borderMethod).RGBA()
			w = (1 - dx) * dy
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			r, g, b, a = getColor(src, x+1, y+1, borderMethod).RGBA()
			w = dx * dy
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			dst.Set(xdst, ydst, color.RGBA64{uint16(R), uint16(G), uint16(B), uint16(A)})
		}
	}
	return dst
}

func bilinearRGBA(src *image.RGBA, transform TransformFunc, width, height, borderMethod int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	b := src.Bounds()
	j := 0
	for ydst := 0; ydst < height; ydst++ {
		for xdst := 0; xdst < width; xdst++ {
			X, Y := transform(xdst, ydst)
			X -= 0.5
			Y -= 0.5
			x := int(X)
			y := int(Y)
			if X < 0.0 {
				x -= 1
			}
			if Y < 0.0 {
				y -= 1
			}
			// Are all neighbours outside the source image?
			if x < -1 || y < -1 || x >= b.Dx() || y >= b.Dy() {
				j += 4
				continue
			}
			// Pixel weights
			dx := X - float64(x-b.Min.X)
			dy := Y - float64(y-b.Min.Y)
			// Image boundaries
			x += b.Min.X
			y += b.Min.Y
			// Calculate new color, add 0.5 to round to nearest integer.
			var R, G, B, A float64 = 0.5, 0.5, 0.5, 0.5
			w := (1 - dx) * (1 - dy)
			r, g, b, a := getRGBA(src, x, y, borderMethod)
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			r, g, b, a = getRGBA(src, x+1, y, borderMethod)
			w = dx * (1 - dy)
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			r, g, b, a = getRGBA(src, x, y+1, borderMethod)
			w = (1 - dx) * dy
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			r, g, b, a = getRGBA(src, x+1, y+1, borderMethod)
			w = dx * dy
			R += w * float64(r)
			G += w * float64(g)
			B += w * float64(b)
			A += w * float64(a)
			dst.Pix[j] = uint8(R)
			j++
			dst.Pix[j] = uint8(G)
			j++
			dst.Pix[j] = uint8(B)
			j++
			dst.Pix[j] = uint8(A)
			j++
		}
	}
	return dst
}

func bilinearGray(src *image.Gray, transform TransformFunc, width, height, borderMethod int) image.Image {
	dst := image.NewGray(image.Rect(0, 0, width, height))
	b := src.Bounds()
	j := 0
	for ydst := 0; ydst < height; ydst++ {
		for xdst := 0; xdst < width; xdst++ {
			X, Y := transform(xdst, ydst)
			X -= 0.5
			Y -= 0.5
			x := int(X)
			y := int(Y)
			if X < 0.0 {
				x -= 1
			}
			if Y < 0.0 {
				y -= 1
			}
			// Are all neighbours outside the source image?
			if x < -1 || y < -1 || x >= b.Dx() || y >= b.Dy() {
				j++
				continue
			}
			// Pixel weights
			dx := X - float64(x)
			dy := Y - float64(y)
			// Image boundaries
			x += b.Min.X
			y += b.Min.Y
			// Calculate new color, add 0.5 to round to nearest integer.
			var G float64 = 0.5
			G += float64(getGray(src, x, y, borderMethod)) * (1 - dx)
			G += float64(getGray(src, x+1, y, borderMethod)) * dx
			G *= (1 - dy)
			var g float64 = 0
			g += float64(getGray(src, x, y+1, borderMethod)) * (1 - dx)
			g += float64(getGray(src, x+1, y+1, borderMethod)) * dx
			G += g * dy
			dst.Pix[j] = uint8(G)
			j++
		}
	}
	return dst
}

func Bicubic(src image.Image, transform TransformFunc, width, height, borderMethod int) image.Image {
	dst := newImage(src, width, height)
	b := src.Bounds()
	for ydst := 0; ydst < height; ydst++ {
		for xdst := 0; xdst < width; xdst++ {
			X, Y := transform(xdst, ydst)
			X -= 0.5
			Y -= 0.5
			x := int(X)
			y := int(Y)
			if X < 0.0 {
				x -= 1
				//dx = 1 - dx
			}
			if Y < 0.0 {
				y -= 1
				//dy = 1 - dy
			}
			// Are all neighbours outside the source image?
			if x < -1 || y < -1 || x >= b.Dx() || y >= b.Dy() {
				//continue
			}
			// Pixel weights
			dx := X - float64(x)
			dy := Y - float64(y)
			// Image boundaries
			x += b.Min.X
			y += b.Min.Y
			// Calculate new color, add 0.5 to round to nearest integer.
			var R, G, B, A float64 = 0.5, 0.5, 0.5, 0.5
			fmt.Println(xdst, ydst, " -> ", X, Y, x, y, dx, dy)
			r00, g00, b00, a00 := getColor(src, x-1, y-1, borderMethod).RGBA()
			r10, g10, b10, a10 := getColor(src, x+0, y-1, borderMethod).RGBA()
			r20, g20, b20, a20 := getColor(src, x+1, y-1, borderMethod).RGBA()
			r30, g30, b30, a30 := getColor(src, x+2, y-1, borderMethod).RGBA()
			r01, g01, b01, a01 := getColor(src, x-1, y+0, borderMethod).RGBA()
			r11, g11, b11, a11 := getColor(src, x+0, y+0, borderMethod).RGBA()
			r21, g21, b21, a21 := getColor(src, x+1, y+0, borderMethod).RGBA()
			r31, g31, b31, a31 := getColor(src, x+2, y+0, borderMethod).RGBA()
			r02, g02, b02, a02 := getColor(src, x-1, y+1, borderMethod).RGBA()
			r12, g12, b12, a12 := getColor(src, x+0, y+1, borderMethod).RGBA()
			r22, g22, b22, a22 := getColor(src, x+1, y+1, borderMethod).RGBA()
			r32, g32, b32, a32 := getColor(src, x+2, y+1, borderMethod).RGBA()
			r03, g03, b03, a03 := getColor(src, x-1, y+2, borderMethod).RGBA()
			r13, g13, b13, a13 := getColor(src, x+0, y+2, borderMethod).RGBA()
			r23, g23, b23, a23 := getColor(src, x+1, y+2, borderMethod).RGBA()
			r33, g33, b33, a33 := getColor(src, x+2, y+2, borderMethod).RGBA()
			fmt.Println(r11, g11, b11, a11)
			r0 := cubicSpline(dx, float64(r00), float64(r10), float64(r20), float64(r30))
			g0 := cubicSpline(dx, float64(g00), float64(g10), float64(g20), float64(g30))
			b0 := cubicSpline(dx, float64(b00), float64(b10), float64(b20), float64(b30))
			a0 := cubicSpline(dx, float64(a00), float64(a10), float64(a20), float64(a30))
			r1 := cubicSpline(dx, float64(r01), float64(r11), float64(r21), float64(r31))
			g1 := cubicSpline(dx, float64(g01), float64(g11), float64(g21), float64(g31))
			b1 := cubicSpline(dx, float64(b01), float64(b11), float64(b21), float64(b31))
			a1 := cubicSpline(dx, float64(a01), float64(a11), float64(a21), float64(a31))
			r2 := cubicSpline(dx, float64(r02), float64(r12), float64(r22), float64(r32))
			g2 := cubicSpline(dx, float64(g02), float64(g12), float64(g22), float64(g32))
			b2 := cubicSpline(dx, float64(b02), float64(b12), float64(b22), float64(b32))
			a2 := cubicSpline(dx, float64(a02), float64(a12), float64(a22), float64(a32))
			r3 := cubicSpline(dx, float64(r03), float64(r13), float64(r23), float64(r33))
			g3 := cubicSpline(dx, float64(g03), float64(g13), float64(g23), float64(g33))
			b3 := cubicSpline(dx, float64(b03), float64(b13), float64(b23), float64(b33))
			a3 := cubicSpline(dx, float64(a03), float64(a13), float64(a23), float64(a33))
			R += cubicSpline(dy, r0, r1, r2, r3)
			G += cubicSpline(dy, g0, g1, g2, g3)
			B += cubicSpline(dy, b0, b1, b2, b3)
			A += cubicSpline(dy, a0, a1, a2, a3)
			R = math.Max(math.Min(R, 65535), 0)
			G = math.Max(math.Min(G, 65535), 0)
			B = math.Max(math.Min(B, 65535), 0)
			A = math.Max(math.Min(A, 65535), 0)
			fmt.Println(R, G, B, A)
			dst.Set(xdst, ydst, color.RGBA64{uint16(R), uint16(G), uint16(B), uint16(A)})
		}
	}
	return dst
}

func cubicSpline(x, p0, p1, p2, p3 float64) float64 {
	return p1 + 0.5*x*(p2-p0+x*(2.0*p0-5.0*p1+4.0*p2-p3+x*(3.0*(p1-p2)+p3-p0)))
}

func getColor(src image.Image, x, y, borderMethod int) color.Color {
	bound := src.Bounds()
	if x < 0 {
		switch borderMethod {
		case BORDER_COPY:
			x = 0
		default:
			return color.RGBA64{}
		}
	} else if x >= bound.Max.X {
		switch borderMethod {
		case BORDER_COPY:
			x = bound.Max.X - 1
		default:
			return color.RGBA64{}
		}
	}
	if y < 0 {
		switch borderMethod {
		case BORDER_COPY:
			y = 0
		default:
			return color.RGBA64{}
		}
	} else if y >= bound.Max.Y {
		switch borderMethod {
		case BORDER_COPY:
			y = bound.Max.Y - 1
		default:
			return color.RGBA64{255, 255, 255, 0}
		}
	}
	return src.At(x, y)
}

func getRGBA(src *image.RGBA, x, y, borderMethod int) (r, g, b, a uint8) {
	bound := src.Bounds()
	if x < 0 {
		switch borderMethod {
		case BORDER_COPY:
			x = 0
		default:
			return 0, 0, 0, 0
		}
	} else if x >= bound.Max.X {
		switch borderMethod {
		case BORDER_COPY:
			x = bound.Max.X - 1
		default:
			return 0, 0, 0, 0
		}
	}
	if y < 0 {
		switch borderMethod {
		case BORDER_COPY:
			y = 0
		default:
			return 0, 0, 0, 0
		}
	} else if y >= bound.Max.Y {
		switch borderMethod {
		case BORDER_COPY:
			y = bound.Max.Y - 1
		default:
			return 0, 0, 0, 0
		}
	}
	i := (y-bound.Min.Y)*src.Stride + (x-bound.Min.X)*4
	return src.Pix[i], src.Pix[i+1], src.Pix[i+2], src.Pix[i+3]
}

func getGray(src *image.Gray, x, y, borderMethod int) uint8 {
	bound := src.Bounds()
	if x < 0 {
		switch borderMethod {
		case BORDER_COPY:
			x = 0
		default:
			return 0
		}
	} else if x >= bound.Max.X {
		switch borderMethod {
		case BORDER_COPY:
			x = bound.Max.X - 1
		default:
			return 0
		}
	}
	if y < 0 {
		switch borderMethod {
		case BORDER_COPY:
			y = 0
		default:
			return 0
		}
	} else if y >= bound.Max.Y {
		switch borderMethod {
		case BORDER_COPY:
			y = bound.Max.Y - 1
		default:
			return 0
		}
	}
	i := (y-bound.Min.Y)*src.Stride + (x - bound.Min.X)
	return src.Pix[i]
}

func newImage(m image.Image, width, height int) WritableImage {
	switch m.(type) {
	case *image.RGBA:
		return image.NewRGBA(image.Rect(0, 0, width, height))
	case *image.Gray:
		return image.NewGray(image.Rect(0, 0, width, height))
	case *image.Gray16:
		return image.NewGray16(image.Rect(0, 0, width, height))
	case *image.NRGBA:
		return image.NewNRGBA(image.Rect(0, 0, width, height))
	case *image.RGBA64:
		return image.NewRGBA64(image.Rect(0, 0, width, height))
	case *image.Alpha:
		return image.NewAlpha(image.Rect(0, 0, width, height))
	case *image.Alpha16:
		return image.NewAlpha16(image.Rect(0, 0, width, height))
	}
	return image.NewRGBA(image.Rect(0, 0, width, height))
}
