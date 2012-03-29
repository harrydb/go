// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "github.com/ziutek/blas"

func (A *Matrix) ScaleBLAS(v float64) {

	// Normal matrices.
	if A.stride == A.width {
		blas.Dscal(len(A.data), v, A.data, 1)
		return
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		blas.Dscal(A.width, v, A.Row(i), 1)
	}
}
