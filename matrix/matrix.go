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

	for i := 0; i < a.height; i++ {
		for j := 0; j < a.width; j++ {
			fmt.Fprint(buffer, a.At(i, j), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}
