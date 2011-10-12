// Package transform provides transformations for images of type image.Image.
//
// The affine transformations do not change the image size, they operate on the
// content only. Implemented affine transformations: translate, zoom, shear and
// rotate.
// 
// Interpolation functions: Nearest neighbor and Bilinear
//
// Example 1, simple affine transform:
// var m image.Image
// // load image
// mNew := transform.Rotate(m, math.Pi/6)
// 
// Example 2, chained affine transforms
// t := transform.NewAffineMatrix()
// xcenter := float64(m.Bounds().Dx()) / 2
// ycenter := float64(m.Bounds().Dy()) / 2
// t.AddZoom(0.25, 0.25, xcenter/2, ycenter)
// t.AddRotation(math.Pi/2, xcenter, ycenter)
// m = transform.Apply(t, m, transform.Bilinear)
//
package transform

import (
	"math"
	"image"
)

type AffineMatrix [9]float64

// NewAffineMatrix returns an identity AffineMatrix.
//
// This object can be used to chain multiple affine transformations.
func NewAffineMatrix() AffineMatrix {
	var a AffineMatrix
	a[0] = 1
	a[4] = 1
	a[8] = 1
	return a
}

// Translate returns image m translated horizontally by tx and vertically by ty.
//
// The interpolation function used is transform.Bilinear.
func Translate(m image.Image, tx, ty float64) image.Image {
	a := translationMatrix(tx, ty)
	return Apply(a, m, Bilinear)
}

// Zoom returns image m zoomed horizontally by sx and vertically by sy.
// Note: this function scales the content, not the image itself. If you want to 
// change the image size, use transform.Scale instead.
//
// The interpolation function used is transform.Bilinear.
func Zoom(m image.Image, sx, sy float64) image.Image {
	b := m.Bounds()
	xcenter := float64(b.Dx()) / 2
	ycenter := float64(b.Dy()) / 2
	a := scaleMatrix(sx, sy, xcenter, ycenter)
	return Apply(a, m, Bilinear)
}

// Shear returns image m sheared horizontally by hx and vertically by hy.
//
// The interpolation function used is transform.Bilinear.
func Shear(m image.Image, hx, hy float64) image.Image {
	b := m.Bounds()
	xcenter := float64(b.Dx()) / 2
	ycenter := float64(b.Dy()) / 2
	a := shearMatrix(hx, hy, xcenter, ycenter)
	return Apply(a, m, Bilinear)
}

// Rotate returns image m rotated around the center by theta.
//
// The interpolation function used is transform.Bilinear.
func Rotate(m image.Image, theta float64) image.Image {
	b := m.Bounds()
	xcenter := float64(b.Dx()) / 2
	ycenter := float64(b.Dy()) / 2
	a := rotationMatrix(theta, xcenter, ycenter)
	return Apply(a, m, Bilinear)
}

// AddTranslation adds horizontal translation by tx and vertical translation by ty to AffineMatrix a.
func (a *AffineMatrix) AddTranslation(tx, ty float64) {
	a.mul(translationMatrix(tx, ty))
}

// AddZoom adds horizontal zoom by sx and vertical zoom by sy to AffineMatrix a.
//
// The zoomed image will be centered around xcenter and ycenter. The center of
// an image.Image m can be found using:
//     xcenter := float64(m.Bounds().Dx()) / 2
//     ycenter := float64(m.Bounds().Dy()) / 2
// This is what transform.Zoom uses.
func (a *AffineMatrix) AddZoom(sx, sy, xcenter, ycenter float64) {
	a.mul(scaleMatrix(sx, sy, xcenter, ycenter))
}

// AddShear adds horizontal shear hx and vertical shear hy to AffineMatrix a.
func (a *AffineMatrix) AddShear(hx, hy, xcenter, ycenter float64) {
	a.mul(shearMatrix(hx, hy, xcenter, ycenter))
}

// AddRotation adds rotation by theta to AffineMatrix a.
func (a *AffineMatrix) AddRotation(theta, xcenter, ycenter float64) {
	a.mul(rotationMatrix(theta, xcenter, ycenter))
}

// Apply applies affine transformation t to image.Image m.
//
// InterpolationFunc f is the interpolation function to be used, thin can be
// transform.Bilinear, transform.Nearest or a custom interpolation function.
func Apply(a AffineMatrix, src image.Image, interpolate InterpolationFunc) image.Image {
	// Create closure over AffineMatrix a
	t := func(x, y int) (float64, float64) {
			X := float64(x) + 0.5
			Y := float64(y) + 0.5
			nx := X*a[0] + Y*a[1] + a[2]
			ny := X*a[3] + Y*a[4] + a[5]
		return nx, ny
	}
	return interpolate(src, t, src.Bounds().Dx(), src.Bounds().Dy(), BORDER_TRANSPARENT)
}

// Normal translation matrix:
//            [1, 0, tx]
// T(tx,ty) = [0, 1, ty]
//            [0, 0, 1]
// 
// Inverse mapping translation matrix:
//             [1, 0, -tx]
// T'(tx,ty) = [0, 1, -ty]
//             [0, 0,  1 ]
func translationMatrix(tx, ty float64) AffineMatrix {
	var t AffineMatrix
	t[0] = 1
	t[2] = -tx
	t[4] = 1
	t[5] = -ty
	t[8] = 1
	return t
}

// Normal zoom (scale) matrix:
//            [sx, 0 , 0]
// S(sx,sy) = [0 , sy, 0]
//            [0 , 0 , 1]
// 
// Inverse mapping zoom (scale) matrix:
//             [1/sx, 0   , 0]
// S'(sx,sy) = [0   , 1/sy, 0]
//             [0   , 0   , 1]
func scaleMatrix(sx, sy, xcenter, ycenter float64) AffineMatrix {
	s := translationMatrix(-xcenter, -ycenter)
	var ss AffineMatrix
	ss[0] = 1 / sx
	ss[4] = 1 / sy
	ss[8] = 1
	s.mul(ss)
	s.mul(translationMatrix(xcenter, ycenter))
	return s
}

// shearMatrix returns an AffineMatrix for shearing around xcenter and ycenter.
//
// Normal shear matrix:
//            [1, hx, 0]
// H(hx,hy) = [hy, 1, 0]
//            [0 , 0, 1]
// 
// Inverse mapping shear matrix:
//             [1  ,-hx, 0]
// H'(hx,hy) = [-hy, 1 , 0] / (1 - hx*hy)
//             [0  , 0 , 1]
func shearMatrix(hx, hy, xcenter, ycenter float64) AffineMatrix {
	h := translationMatrix(-xcenter, -ycenter)
	var hh AffineMatrix
	// use inverse mapping
	denom := 1 - (hx * hy)
	hh[0] = 1 / denom
	hh[1] = -hx / denom
	hh[3] = -hy / denom
	hh[4] = 1 / denom
	hh[8] = 1
	h.mul(hh)
	h.mul(translationMatrix(xcenter, ycenter))
	return h
}

// rotationMatrix returns an AffineMatrix for rotation around xcenter and ycenter.
//
// Normal rotation matrix:
//            [cos(theta), sin(theta), 0]
// R(theta) = [-sin(theta), cos(theta), 0]
//            [0         , 0          , 1]
// 
// Inverse mapping rotation matrix:
// R'(theta) = R(-theta)
func rotationMatrix(theta, xcenter, ycenter float64) AffineMatrix {
	r := translationMatrix(-xcenter, -ycenter)
	var rr AffineMatrix
	// use inverse mapping
	rr[0] = math.Cos(theta) // cos(-theta) == cos(theta)
	rr[1] = math.Sin(theta) // -sin(-theta) == sin(theta)
	rr[3] = math.Sin(-theta)
	rr[4] = math.Cos(theta) // cos(-theta) == cos(theta)
	rr[8] = 1
	r.mul(rr)
	r.mul(translationMatrix(xcenter, ycenter))
	return r
}

// mul multiplies AffineMatrix a with AffineMatrix b (storing the result in a).
func (a *AffineMatrix) mul(b AffineMatrix) {
	var c AffineMatrix
	c[0] = a[0]*b[0] + a[1]*b[3] + a[2]*b[6]
	c[1] = a[0]*b[1] + a[1]*b[4] + a[2]*b[7]
	c[2] = a[0]*b[2] + a[1]*b[5] + a[2]*b[8]
	c[3] = a[3]*b[0] + a[4]*b[3] + a[5]*b[6]
	c[4] = a[3]*b[1] + a[4]*b[4] + a[5]*b[7]
	c[5] = a[3]*b[2] + a[4]*b[5] + a[5]*b[8]
	c[6] = a[6]*b[0] + a[7]*b[3] + a[8]*b[6]
	c[7] = a[6]*b[1] + a[7]*b[4] + a[8]*b[7]
	c[8] = a[6]*b[2] + a[7]*b[5] + a[8]*b[8]
	*a = c
}

// invertMatrix returns the inverse of AffineMatrix a.
func (a *AffineMatrix) invertMatrix() AffineMatrix {
	var b AffineMatrix
	A := a[4]*a[8] - a[5]*a[7]
	B := a[5]*a[6] - a[3]*a[8]
	C := a[3]*a[7] - a[4]*a[6]
	det := a[0]*A + a[1]*B + a[2]*C
	b[0] = A / det
	b[3] = B / det
	b[6] = C / det
	b[1] = (a[2]*a[7] - a[1]*a[8]) / det
	b[4] = (a[0]*a[8] - a[2]*a[6]) / det
	b[7] = (a[6]*a[1] - a[0]*a[7]) / det
	b[2] = (a[1]*a[5] - a[2]*a[4]) / det
	b[5] = (a[2]*a[3] - a[0]*a[5]) / det
	b[8] = (a[0]*a[4] - a[1]*a[3]) / det
	return b
}
