package main

import (
	"os"
	"log"
	//"time"
	//"image"
	"image/pnm"
	//"image/transform"
	"image/gray"
)

func main() {
	//file, err := os.Open("Aug-0011.ppm")
	file, err := os.Open("big.ppm")
	//file, err := os.Open("Aug-0011.ppm")
	//file, err := os.Open("test.pgm")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Decoding...")
	img, err := pnm.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("To gray...")
	grayImg := gray.Convert(img, gray.ToGrayLuma709)
	//log.Println("Otsu...")
	//threshold := gray.Otsu(grayImg)
	//log.Println("Thresholding...")
	//gray.Threshold(grayImg, threshold, 0, 255)
	//log.Println("Cocos...")
	//tic := time.Nanoseconds()
	//cocos := gray.CoCos(grayImg, 255)
	//toc := time.Nanoseconds() - tic
	//log.Printf("time: %.3fs\n", float32(toc)/1000000000)
	//log.Println(len(cocos))
//
	//grayImg = transform.Shear(grayImg, 1.1, 0).(*image.Gray)
	//b := grayImg.Bounds()
	//crop := 0.55 * float64(b.Dy())
	//b.Min.X += int(crop)
	//b.Max.X -= int(crop)
	//grayImg = grayImg.SubImage(b).(*image.Gray)
//

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
