// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"bytes"
	"fmt"
)

type Matrix struct {
	height, width int
	stride        int
	data          []float64
}

func (a *Matrix) String() string {
	buffer := bytes.NewBufferString("")

	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			fmt.Fprint(buffer, a.At(x, y), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}
