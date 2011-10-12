package main

import (
	"os"
	"fmt"
	"math"
	"time"
	"image"
	 "image/png"
	_ "image/jpeg"
	"image/transform"
)

func main() {
	//r, err := os.Open("IMG_1549.PNG")
	r, err := os.Open("test.png")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println("Opening...")
	m, _, err := image.Decode(r)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println("Transforming...")
	tic := time.Nanoseconds()
	//m = transform.ScaleFactor(m, 3, 3, transform.Bilinear)
	//m = transform.ScaleFactor(m, 0.25, 0.25, transform.Bilinear)
	//m = transform.Shear(m, 1.0, 0)
	//m = transform.Rotate(m, math.Pi/2)
	//m = transform.Zoom(m, 3, 3)
	//m = transform.Translate(m, 1, 1)
	
	t := transform.NewAffineMatrix()
	//xcenter := float64(m.Bounds().Dx()) / 2
	//ycenter := float64(m.Bounds().Dy()) / 2
	//t.AddZoom(0.25, 0.25, xcenter, ycenter)
	t.AddRotation(-math.Pi/6, 3.5, 3.5)
	m = transform.Apply(t, m, transform.Bicubic)
	toc := time.Nanoseconds() - tic
	fmt.Printf("time: %.3fs\n", float32(toc)/1000000000)

	fmt.Println("Writing...")
	w, err := os.Create("out.png")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	err = png.Encode(w, m)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
