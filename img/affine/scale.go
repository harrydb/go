// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package affine

import (
	"image"
)

// Scale returns a copy of the original image scaled to the specified width and
// height. Scale does change the canvas size.
func Scale(src image.Image, width, height int, interpolate InterpolationFunc) image.Image {
	b := src.Bounds()
	wx := float64(b.Dx()) / float64(width)
	wy := float64(b.Dy()) / float64(height)
	// Create closure over wx and wy
	t := func(x, y int) (float64, float64) {
		nx := wx * (0.5 + float64(x))
		ny := wy * (0.5 + float64(y))
		return nx, ny
	}
	return interpolate(src, t, width, height, BORDER_COPY)
}

// ScaleFactor returns a copy of the original image scaled by the specified sx
// and sy. ScaleFactor does change the canvas size.
func ScaleFactor(src image.Image, sx, sy float64, interpolate InterpolationFunc) image.Image {
	b := src.Bounds()
	width := int(float64(b.Dx())*sx + 0.5)
	height := int(float64(b.Dy())*sy + 0.5)
	return Scale(src, width, height, interpolate)
}
