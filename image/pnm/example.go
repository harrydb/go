package main

import (
	"os"
	"log"
	"image/pnm"
)

func main() {
	file, err := os.Open("in.ppm")
	if err != nil {
		log.Fatal(err)
	}

	img, err := pnm.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create("out.pnm")
	if err != nil {
		log.Fatal(err)
	}

	err = pnm.Encode(outFile, img, pnm.PPM)
	if err != nil {
		log.Fatal(err)
	}

}
