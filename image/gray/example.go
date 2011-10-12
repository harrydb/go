package main

import (
	"os"
	"log"
	"image/pnm"
	"image/gray"
)

func main() {
	file, err := os.Open("in.ppm")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Decoding...")
	img, err := pnm.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("To gray...")
	grayImg := gray.Convert(img, gray.ToGrayLuminance)
	log.Println("Otsu...")
	threshold := gray.Otsu(grayImg)
	log.Println("Thresholding...")
	gray.Threshold(grayImg, threshold, 0, 255)
	log.Println("Cocos...")
	cocos := gray.CoCos(grayImg, 255)
	log.Println(len(cocos))

	log.Println("Writing...")
	outFile, err := os.Create("out.pnm")
	if err != nil {
		log.Fatal(err)
	}

	err = pnm.Encode(outFile, grayImg, pnm.PGM)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done...")

}
