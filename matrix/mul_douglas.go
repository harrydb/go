// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulDouglas returns A * B.
//
// This function implements the Strassen-Winograd matrix multiplication
// algorithm with additional memory usage < 2/3(n^2) for n x n matrices.
// For non-square and odd-sized matrices it falls back to naive matrix
// multiplication. When n < it also falls back.
//
//	Original paper:
//	Douglas et al, 1994.
//	GEMMW: A Portable Level 3 BLAS Winograd Variant of Strassen's Matrix-Matrix Multiply Algorithm.
//	http://www.cs.yale.edu/publications/techreports/tr904.pdf
func MulDouglas(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulDouglas(A, B)
}

func (C *Matrix) MulDouglas(A, B *Matrix) *Matrix {

	if A.width < 80 || A.height % 2 != 0 || A.width % 2 != 0 || B.width % 2 != 0 {
		return C.MulBLAS(A, B)
	}

	m := A.height / 2
	k := A.width / 2
	n := B.width / 2
	A11 := A.SubMatrix(0, 0, m, k)
	A12 := A.SubMatrix(0, n, m, k)
	A21 := A.SubMatrix(n, 0, m, k)
	A22 := A.SubMatrix(n, n, m, k)
	B11 := B.SubMatrix(0, 0, n, n)
	B12 := B.SubMatrix(0, n, n, n)
	B21 := B.SubMatrix(n, 0, n, n)
	B22 := B.SubMatrix(n, n, n, n)
	C11 := C.SubMatrix(0, 0, n, n)
	C12 := C.SubMatrix(0, n, n, n)
	C21 := C.SubMatrix(n, 0, n, n)
	C22 := C.SubMatrix(n, n, n, n)

	// Allocate scratch space
	X := Zeros(n, n)
	Y := Zeros(n, n)

	// Perform calculations.
	X.Minus(A11, A21)
	Y.Minus(B22, B12)
	C21.MulDouglas(X, Y)
	X.PlusBLAS(A21, A22)
	Y.Minus(B12, B11)
	C22.MulDouglas(X, Y)
	X.SubBLAS(A11)
	Y.Minus(B22, Y)
	C12.MulDouglas(X, Y)
	X.Minus(A12, X)
	C11.MulDouglas(X, B22)
	X.MulDouglas(A11, B11)
	C12.AddBLAS(X)
	C21.AddBLAS(C12)
	C12.AddBLAS(C22)
	C22.AddBLAS(C21) // Final c22.
	C12.AddBLAS(C11) // Final c12.
	Y.SubBLAS(B21)
	C11.MulDouglas(A22, Y)
	C21.SubBLAS(C11) // Final c21.
	C11.MulDouglas(A12, B21)
	C11.AddBLAS(X) // Final c11.
	return C
}
