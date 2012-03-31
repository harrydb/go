// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"testing"
	"math"
	"math/rand"
)

func TestSubMatrix(t *testing.T) {
	a := randomMatrix(4, 4)
	c := a.SubMatrix(1, 1, 2, 2)

	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			if a.At(i + 1, j + 1) != c.At(i, j) {
				t.Fatalf("Element a(%d,%d) should be the same as b(%d, %d)", i, j, j, i)
			}
		}
	}
}

func TestIdentity(t *testing.T) {
	n := 10
	I := Identity(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				if I.At(i, j) != 1.0 {
					t.Fatalf("Element at (%d,%d) should be 1.0", i, j)
				}
			} else {
				if I.At(i, j) != 0.0 {
					t.Fatalf("Element at (%d,%d) should be 0.0", i, j)
				}
			}
		}
	}

}

func TestTranspose(t *testing.T) {
	m := 2
	n := 3
	A := randomMatrix(m, n)
	B := Transpose(A)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A.At(i, j) != B.At(j, i) {
				t.Fatalf("Element A(%d,%d) should be the same as B(%d, %d)", i, j, j, i)
			}
		}
	}
}

func randomMatrix(m, n int) (A *Matrix) {
	A = Zeros(m, n)
	A.randomize()
	return A
}

func (A *Matrix) randomize() {
	for i := range A.data {
		A.data[i] = rand.Float64()
	}
}

func equal(A, B *Matrix, ε float64, t *testing.T) bool {
	if A.height != B.height || A.width != B.width {
		t.Fatalf("Wrong result: different size A: %d x %d \n A: %d x %d \n", A.height, A.width, B.height, B.width)
	}
	for i := 0; i < A.height; i++ {
		for j := 0; j < A.width; j++ {
			if math.Abs(A.At(i, j) - B.At(i,j)) > ε {
				t.Logf("Wrong result: A(%d, %d) = %v \t B(%d, %d) = %v \n", i, j, A.At(i, j), i, j, B.At(i, j))
				return false
			}
		}
	}
	return true
}

