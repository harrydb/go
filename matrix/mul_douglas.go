// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulDouglas returns A * B.
// See Douglas et al, 1994.
// http://www.cs.yale.edu/publications/techreports/tr904.pdf
func MulDouglas(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulAddDouglas(A, B)
}

func (C *Matrix) MulAddDouglas(A, B *Matrix) *Matrix {

	if A.width < 80 || A.height != A.width || A.height % 2 != 0 {
		C.MulAddBLAS(A, B)
		return C
	}

	m := A.height / 2
	A11 := A.SubMatrix(0, 0, m, m)
	A12 := A.SubMatrix(0, m, m, m)
	A21 := A.SubMatrix(m, 0, m, m)
	A22 := A.SubMatrix(m, m, m, m)
	B11 := B.SubMatrix(0, 0, m, m)
	B12 := B.SubMatrix(0, m, m, m)
	B21 := B.SubMatrix(m, 0, m, m)
	B22 := B.SubMatrix(m, m, m, m)
	C11 := C.SubMatrix(0, 0, m, m)
	C12 := C.SubMatrix(0, m, m, m)
	C21 := C.SubMatrix(m, 0, m, m)
	C22 := C.SubMatrix(m, m, m, m)

	X := Minus(A11, A21) // Allocate scratch space 1.
	Y := Minus(B22, B12) // Allocate scratch space 2.
	C21.MulAddDouglas(X, Y)
	X.Plus(A21, A22)
	Y.Minus(B12, B11)
	C22.MulAddDouglas(X, Y)
	X.Sub(A11)
	Y.Minus(B22, Y)
	C12.MulAddDouglas(X, Y)
	X.Minus(A12, X)
	C11.MulAddDouglas(X, B22)
	X.Clear() // Clear scratch for reuse.
	X.MulAddDouglas(A11, B11)
	C12.Add(X)
	C21.Add(C12)
	C12.Add(C22)
	C22.Add(C21) // Final c22.
	C12.Add(C11) // Final c12.
	Y.Sub(B21)
	C11.Clear() // Clear scratch for reuse.
	C11.MulAddDouglas(A22, Y)
	C21.Sub(C11) // Final c21.
	C11.Clear() // Clear scratch for reuse.
	C11.MulAddDouglas(A12, B21)
	C11.Add(X) // Final c11.
	return C
}
