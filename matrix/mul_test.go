// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math"
	"math/rand"
	"testing"
	"time"
	//gomatrix "code.google.com/p/gomatrix/matrix"
	gomatrix "harrydb1984-gomatrix/matrix"
)

const ε = 10e-12

func TestMulNaive(t *testing.T) {
	A := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	B := New(3, 2, []float64{8, 4, 2, 2, 1, 0})
	C := MulNaive(A, B)
	D := New(2, 2, []float64{16, 8, 20, 10})

	if !equal(C, D, 0, t) {
		t.FailNow()
	}
}

func TestMulBLAS(t *testing.T) {
	n := 200
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := MulNaive(A, B)
	D := MulBLAS(A, B)

	if !equal(C, D, 0, t) {
		t.FailNow()
	}
}

func TestMulStrassen(t *testing.T) {
	n := 200
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := MulNaive(A, B)
	D := MulStrassen(A, B)

	if !equal(C, D, ε, t) {
		t.FailNow()
	}
}

func TestMulStrassenPar(t *testing.T) {
	n := 200
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := MulStrassen(A, B)
	D := MulStrassenPar(A, B)

	if !equal(C, D, ε, t) {
		t.FailNow()
	}
}

func TestMulWinograd(t *testing.T) {
	rand.Seed(1)
	n := 200
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := MulStrassen(A, B)
	D := MulWinograd(A, B)

	if !equal(C, D, ε, t) {
		t.FailNow()
	}
}

func TestMulDouglas(t *testing.T) {
	n := 200
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := MulNaive(A, B)
	D := MulDouglas(A, B)

	if !equal(C, D, ε, t) {
		t.FailNow()
	}
}

func TestMulHuss(t *testing.T) {
	n := 200
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	C := MulNaive(A, B)
	D := MulHuss(A, B)

	if !equal(C, D, ε, t) {
		t.FailNow()
	}
}

func TestGomatrix(t *testing.T) {
	n := 200
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	C := MulNaive(a, b)
	D, _ := A.TimesDense(B)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if math.Abs(C.At(i, j) - D.Get(i, j)) > ε {
				t.Fatalf("Wrong result: C(%d, %d) = %v \t D(%d, %d) = %v \n", i, j, C.At(i, j), i, j, D.Get(i, j))
			}
		}
	}
}

func BenchmarkMulDouglas__1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulDouglas(A, B)
    }
}

func BenchmarkMulHuss_____1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulHuss(A, B)
    }
}

func BenchmarkMulWinograd_1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulWinograd(A, B)
    }
}

func BenchmarkMulStrassen_1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassen(A, B)
    }
}

func BenchmarkMulStrasPar_1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassenPar(A, B)
    }
}

func BenchmarkMulGomatrix_1024(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 1024
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		A.TimesDense(B)
    }
}

func BenchmarkMulDouglas___512(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 512
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulDouglas(A, B)
    }
}

func BenchmarkMulStrassPar_512(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 512
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassenPar(A, B)
    }
}

func BenchmarkMulStrassen__512(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 512
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassen(A, B)
    }
}

func BenchmarkMulGomatrix__512(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 512
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		A.TimesDense(B)
    }
}

func BenchmarkMulBLAS______512(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 512
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulBLAS(A, B)
    }
}

func BenchmarkMulNaive____512(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 512
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulNaive(A, B)
    }
}

func BenchmarkMulDouglas___256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulDouglas(A, B)
    }
}

func BenchmarkMulHuss______256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulHuss(A, B)
    }
}

func BenchmarkMulStrassPar_256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassenPar(A, B)
    }
}

func BenchmarkMulStrassen__256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassen(A, B)
    }
}

func BenchmarkMulBLAS______256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulBLAS(A, B)
    }
}

func BenchmarkMulGomatrix__256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		A.TimesDense(B)
    }
}

func BenchmarkMulNaive____256(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 256
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulNaive(A, B)
    }
}

func BenchmarkMulStrassPar_128(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 128
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassenPar(A, B)
    }
}

func BenchmarkMulDouglas___128(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 128
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulDouglas(A, B)
    }
}

func BenchmarkMulStrassen__128(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 128
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulStrassen(A, B)
    }
}

func BenchmarkMulBLAS______128(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 128
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulBLAS(A, B)
    }
}

func BenchmarkMulGomatrix__128(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 128
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		A.TimesDense(B)
    }
}

func BenchmarkMulNaive____128(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 128
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulNaive(A, B)
	}
}

func BenchmarkMulBLAS_______64(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulBLAS(A, B)
    }
}

func BenchmarkMulGomatrix___64(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		A.TimesDense(B)
    }
}

func BenchmarkMulNaive_____64(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 64
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulNaive(A, B)
	}
}

func BenchmarkMulBLAS_______32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulBLAS(A, B)
    }
}

func BenchmarkMulGomatrix___32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	a := randomMatrix(n, n)
	A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
	b := randomMatrix(n, n)
	B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		A.TimesDense(B)
    }
}

func BenchmarkMulNaive_____32(bench *testing.B) {
	bench.StopTimer()
	rand.Seed(time.Now().Unix())
	n := 32
	A := randomMatrix(n, n)
	B := randomMatrix(n, n)
	bench.StartTimer()

	for i := 0; i < bench.N; i++ {
		MulNaive(A, B)
	}
}
