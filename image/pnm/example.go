package main

import (
	"os"
	"log"
	"image/pnm"
)

func main() {
	file, err := os.Open("Amsterdam-0002.ppm")
	//file, err := os.Open("a.ppm")
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

	err = pnm.Encode(outFile, img, pnm.PBM)
	if err != nil {
		log.Fatal(err)
	}

}
