// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

// MulWinograd returns C = A * B.
//
// This function implements the Strassen-Winograd variant of the Strassen
// matrix multiplication algorithm. It uses less additions than Strassen's
// original algorithm.
//
// This implementation is not (really) optimized for efficient memory placement
// and usage. Use MulDouglas() instead.
//
//	Original paper:
//	S. Winograd, 1971.
//	On multiplication of 2×2 matrices.
//	http://dx.doi.org/10.1016/0024-3795(71)90009-7
func MulWinograd(A, B *Matrix) *Matrix {
	return Zeros(A.height, B.width).MulWinograd(A, B)
}

// MulWinograd calculates C = A * B and returns C.
func (C *Matrix) MulWinograd(A, B *Matrix) *Matrix {

	if A.width < 80 || A.height != A.width || A.height % 2 != 0 {
		return C.MulBLAS(A, B)
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

	// 8 additions + 8 allocations.
	S1 := Plus(A21, A22)
	S2 := Minus(S1, A11)
	S3 := Minus(A11, A21)
	S4 := Minus(A12, S2)
	T1 := Minus(B12, B11)
	T2 := Minus(B22, T1)
	T3 := Minus(B22, B12)
	T4 := Minus(B21, T2)

	// 7 multiplications
	C22.MulWinograd(S1, T1)
	S1.MulWinograd(A11, B11)
	C11.MulWinograd(A12, B21)

	C21.MulWinograd(A22, T4)
	T4.MulWinograd(S2, T2)
	S2.MulWinograd(S3, T3)
	C12.MulWinograd(S4, B22)

	// 7 additions
	C11.AddBLAS(S1)
	S1.AddBLAS(T4)
	C12.AddBLAS(S1).AddBLAS(C22)
	S1.AddBLAS(S2)
	C21.AddBLAS(S1)
	C22.AddBLAS(S1)
	return C
}
	// // 8 additions + 8 allocations.
	//S1 := Plus(A21, A22)
	//S2 := Minus(S1, A11)
	//S3 := Minus(A11, A21)
	//S4 := Minus(A12, S2)
	//T1 := Minus(B12, B11)
	//T2 := Minus(B22, T1)
	//T3 := Minus(B22, B12)
	//T4 := Minus(B21, T2)
//
	// // 7 multiplications + 7 allocations.
	//P1 := MulWinograd(A11, B11)
	//P2 := MulWinograd(A12, B21)
	//P3 := MulWinograd(S1, T1)
	//P4 := MulWinograd(S2, T2)
	//P5 := MulWinograd(S3, T3)
	//P6 := MulWinograd(S4, B22)
	//P7 := MulWinograd(A22, T4)
//
	// // 7 additions, no extra space is allocated here.
	//C11.Add(P1).Add(P2)
	//U2 := P1.Add(P4)
	//C12.Add(U2).Add(P3).Add(P6)
	//U4 := U2.Add(P5)
	//C21.Add(P7).Add(U4)
	//C22.Add(U4).Add(P3)
