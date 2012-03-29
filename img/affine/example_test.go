// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package affine_test

import (
	"affine"
	"image/png"
	"log"
	"math"
	"os"
)

// Single affine transform.
func Example() {
	file, err := os.Open("in.png")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Decoding...")
	img, err := pnm.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// A simple rotation.
	outImg := affine.Rotate(img, math.Pi/6)

	log.Println("Writing...")
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outFile, outImg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done...")

}

// By multiplying affine transformation matrices, more than one transformation
// can be performed at once. Convenience functions for this are used in the
// example below.
func Example() {
	file, err := os.Open("in.png")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Decoding...")
	img, err := pnm.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create new affine transformation matrix.
	t := affine.NewAffineMatrix()
	xcenter := float64(outImg.Bounds().Dx()) / 2
	ycenter := float64(outImg.Bounds().Dy()) / 2

	// Modify the transformation matrix:
	t.AddZoom(0.25, 0.25, xcenter/2, ycenter)
	t.AddRotation(math.Pi/2, xcenter, ycenter)

	// Apply the transformation.
	log.Println("Transforming...")
	outImg = affine.Apply(t, outImg, affine.Bilinear)

	log.Println("Writing...")
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outFile, outImg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done...")

}
