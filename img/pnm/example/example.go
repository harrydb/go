// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/harrydb/go/img/pnm"
	"image"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("usage:", os.Args[0], "input-file output-file")
	}
	in, out := os.Args[1], os.Args[2]
	f, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	log.Println("Format:", format)

	f, err = os.Create(out)
	if err != nil {
		log.Fatal(err)
	}

	err = pnm.Encode(f, img, pnm.PPM)
	if err != nil {
		os.Remove(out)
		log.Fatal(err)
	}
	f.Close() // in some situations, this could be important for flushing remaining output.
}
