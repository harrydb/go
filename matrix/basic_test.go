// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"testing"
	"math/rand"
)

// Copy creates a new matrix with identical content.
func Copy(A *Matrix) *Matrix {
	B := Zeros(A.height, A.width)
	B.Add(A)
	return B
}

func TestSubMatrix(t *testing.T) {
	a := randomMatrix(4, 4)
	c := a.SubMatrix(1, 1, 2, 2)

	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			if a.At(j + 1, i + 1) != c.At(j, i) {
				t.Fatalf("Element a(%d,%d) should be the same as b(%d, %d)", j, i, i, j)
			}
		}
	}
}

func TestAdd(t *testing.T) {
	n := 4
	a := randomMatrix(n, n)
	b := randomMatrix(n, n)
	c := a.SubMatrix(2, 2, 2, 2)
	d := b.SubMatrix(2, 2, 2, 2)
	c.Add(d)
}

func TestIdentity(t *testing.T) {
	n := 10
	I := Identity(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				if I.At(j, i) != 1.0 {
					t.Fatalf("Element at (%d,%d) should be 1.0", j, i)
				}
			} else {
				if I.At(j, i) != 0.0 {
					t.Fatalf("Element at (%d,%d) should be 0.0", j, i)
				}
			}
		}
	}

}

func TestTranspose(t *testing.T) {
	m := 2
	n := 3
	a := randomMatrix(m, n)
	b := Transpose(a)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a.At(j, i) != b.At(i, j) {
				t.Fatalf("Element a(%d,%d) should be the same as b(%d, %d)", j, i, i, j)
			}
		}
	}
}

func randomMatrix(m, n int) (a *Matrix) {
	a = Zeros(m, n)

	for i := range a.data {
		a.data[i] = rand.Float64()
	}

	return a
}
