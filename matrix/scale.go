// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

func (A *Matrix) Scale(v float64) {

	// Normal matrices.
	if A.stride == A.width {
		for i, ai := range A.data {
			A.data[i] = ai * v
		}
		return
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		Ai := A.Row(i)
		for j, aij := range Ai {
			Ai[j] = aij * v
		}
	}

}
