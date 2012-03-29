// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package grayscale provides conversion to grayscale images, thresholding and connected component labelling.
//
// Example: converting an image to grayscale
//	grayImg := grayscale.Convert(colorImg, grayscale.ToGrayLuminance)
//
// Example: thresholding an image
//	threshold := gray.Otsu(grayImg)
//	grayscale.Threshold(grayImg, threshold, 0, 255)
//
// Example: find all white connected components
//	cocos := grayscale.CoCos(grayImg, 255)
package grayscale
