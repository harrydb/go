// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

func MulStrassenPar(A, B * Matrix) *Matrix {
	return Zeros(A.height, A.width).MulAddStrassenPar(A, B)
}

func (C *Matrix) MulAddStrassenPar(A, B *Matrix) *Matrix {

	if A.width < 200 || A.height != A.width || A.height % 2 != 0 {
		return C.MulDouglas(A, B)
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

	var M1, M2, M3 *Matrix
	done1 := make(chan int)
	done2 := make(chan int)

	go func() {
		M1 = Zeros(m, m)
		M2 = Zeros(m, m)
		M3 = Zeros(m, m)
		M1.MulAddStrassenPar(Plus(A11, A22), Plus(B11, B22))
		M2.MulAddStrassenPar(Plus(A21, A22), B11)
		M3.MulAddStrassenPar(A11, Minus(B12, B22))
		done1 <- 1
	}()

	go func() {
		C21.MulAddStrassenPar(A22, Minus(B21, B11))
		C12.MulAddStrassenPar(Plus(A11, A12), B22)
		C22.MulAddStrassenPar(Minus(A21, A11), Plus(B11, B12))
		C11.MulAddStrassenPar(Minus(A12, A22), Plus(B21, B22))
		done2 <- 1
	}()

	// Wait for goroutines to finish.
	<- done1
	<- done2

	C11.AddBLAS(M1).AddBLAS(C21).SubBLAS(C12)
	C12.AddBLAS(M3)
	C21.AddBLAS(M2)
	C22.AddBLAS(M1).SubBLAS(M2).AddBLAS(M3)
	return C
}
