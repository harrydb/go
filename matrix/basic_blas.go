// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "github.com/ziutek/blas"

// Copy the contents of B to A.
func (A *Matrix) CopyBLAS(B *Matrix) {

	// Normal matrices.
	if B.stride == B.width {
		blas.Dcopy(len(A.data), B.data, 1, A.data, 1)
		return
	}

	// Submatrices.
	for i := 0; i < A.height; i++ {
		blas.Dcopy(A.width, B.Row(i), 1, A.Row(i), 1)
	}
}
