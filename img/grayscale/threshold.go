// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grayscale

import (
	"image"
)

// Threshold performs thresholding on a grayscale image.
//
// All color values greater than threshold are set to fgColor, the other to bgColor.
func Threshold(m *image.Gray, threshold, bgColor, fgColor uint8) {
	count := len(m.Pix)
	for i := 0; i < count; i++ {
		if m.Pix[i] > threshold {
			m.Pix[i] = fgColor
		} else {
			m.Pix[i] = bgColor
		}
	}
}

// Otsu determines a threshold value using Otsu's method on grayscale images.
//
// "A threshold selection method from gray-level histograms" (Otsu, 1979).
func Otsu(m *image.Gray) uint8 {
	hist := Histogram(m)
	sum := 0
	for i, v := range hist {
		sum += i * v
	}
	wB, wF := 0, len(m.Pix)
	sumB, sumF := 0, sum
	maxVariance := 0.0
	threshold := uint8(0)
	for t := 0; t < 256; t++ {
		wB += hist[t]
		wF -= hist[t]
		if wB == 0 {
			continue
		}
		if wF == 0 {
			return threshold
		}
		sumB += t * hist[t]
		sumF = sum - sumB
		mB := float64(sumB) / float64(wB)
		mF := float64(sumF) / float64(wF)
		betweenVariance := float64(wB*wF) * (mB - mF) * (mB - mF)
		if betweenVariance > maxVariance {
			maxVariance = betweenVariance
			threshold = uint8(t)
		}
	}
	return threshold
}

// Histogram creates a histogram of a grayscale image.
func Histogram(m *image.Gray) []int {
	hist := make([]int, 256)
	count := len(m.Pix)
	for i := 0; i < count; i++ {
		hist[m.Pix[i]]++
	}
	return hist
}
