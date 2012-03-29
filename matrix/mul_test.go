// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math/rand"
	"testing"
	"time"
	"math"
	gomatrix "code.google.com/p/gomatrix/matrix"
)

func TestMulSimple(t *testing.T) {
	a := New(2, 3, []float64{1, 2, 4, 2, 1, 2})
	b := New(3, 2, []float64{8, 4, 2, 2, 1, 0})
	c := MulSimple(a, b)
	correct := New(2, 2, []float64{16, 8, 20, 10})
	if c.height != correct.height || c.width != correct.width {
		t.Fatalf("Wrong result:\n %v Should be:\n %v", c, correct)
	}
	for i := 0; i < correct.height; i++ {
		for j := 0; j < correct.width; j++ {
			if c.At(i, j) != correct.At(i,j) {
				t.Fatalf("Wrong result:\n %v Should be:\n %v", c, correct)
			}
		}
	}
}

func TestMulBLAS(t *testing.T) {
	n := 128
	a := randomMatrix(128, 128)
	b := randomMatrix(128, 128)
	c := MulSimple(a, b)
	d := MulBLAS(a, b)
	ε := 10e-12

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			e := math.Abs(c.At(i, j) - d.At(i, j))
			if e > ε {
				t.Fatalf("Wrong result: %v Should be %v", d.At(i, j), c.At(i, j))
			}
		}
	}
}

func TestMulStrassen(t *testing.T) {
	n := 128
	a := randomMatrix(128, 128)
	b := randomMatrix(128, 128)
	c := MulBLAS(a, b)
	d := MulStrassen(a, b)
	ε := 10e-12

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			e := math.Abs(c.At(i, j) - d.At(i, j))
			if e > ε {
				t.Fatalf("Wrong result: %v Should be %v", d.At(i, j), c.At(i, j))
			}
		}
	}
}

func TestMulStrassenPar(t *testing.T) {
	n := 128
	a := randomMatrix(128, 128)
	b := randomMatrix(128, 128)
	c := MulStrassen(a, b)
	d := MulStrassenPar(a, b)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			e := math.Abs(c.At(i, j) - d.At(i, j))
			if e != 0 {
				t.Fatalf("Wrong result: %v Should be %v", d.At(i, j), c.At(i, j))
			}
		}
	}
}

func TestMulDouglas(t *testing.T) {
	n := 128
	a := randomMatrix(128, 128)
	b := randomMatrix(128, 128)
	c := MulBLAS(a, b)
	d := MulDouglas(a, b)
	ε := 10e-12

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			e := math.Abs(c.At(i, j) - d.At(i, j))
			if e > ε {
				t.Fatalf("Wrong result: (%d, %d) = %v Should be %v", i, j, d.At(i, j), c.At(i, j))
			}
		}
	}
}

func BenchmarkMulDouglas__1024(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(1024, 1024)
		b := randomMatrix(1024, 1024)
		bench.StartTimer()
		MulDouglas(a, b)
    }
}

func BenchmarkMulStrassPar1024(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(1024, 1024)
		b := randomMatrix(1024, 1024)
		bench.StartTimer()
		MulStrassenPar(a, b)
    }
}

func BenchmarkMulDouglas__512(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(512, 512)
		b := randomMatrix(512, 512)
		bench.StartTimer()
		MulDouglas(a, b)
    }
}

func BenchmarkMulStrassPar512(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(512, 512)
		b := randomMatrix(512, 512)
		bench.StartTimer()
		MulStrassenPar(a, b)
    }
}

func BenchmarkMulStrassen_512(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(512, 512)
		b := randomMatrix(512, 512)
		bench.StartTimer()
		MulStrassen(a, b)
    }
}


func BenchmarkMulBLAS_____512(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(512, 512)
		b := randomMatrix(512, 512)
		bench.StartTimer()
		MulBLAS(a, b)
    }
}

func BenchmarkMulSimple___512(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(512, 512)
		b := randomMatrix(512, 512)
		bench.StartTimer()
		MulSimple(a, b)
    }
}

func BenchmarkMulGomatrix_512(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(512, 512)
		A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
		b := randomMatrix(512, 512)
		B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
		bench.StartTimer()
		A.TimesDense(B)
    }
}

func BenchmarkMulDouglas__256(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(256, 256)
		b := randomMatrix(256, 256)
		bench.StartTimer()
		MulDouglas(a, b)
    }
}

func BenchmarkMulStrassPar256(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(256, 256)
		b := randomMatrix(256, 256)
		bench.StartTimer()
		MulStrassenPar(a, b)
    }
}

func BenchmarkMulStrassen_256(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(256, 256)
		b := randomMatrix(256, 256)
		bench.StartTimer()
		MulStrassen(a, b)
    }
}

func BenchmarkMulBLAS_____256(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(256, 256)
		b := randomMatrix(256, 256)
		bench.StartTimer()
		MulBLAS(a, b)
    }
}

func BenchmarkMulSimple___256(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(256, 256)
		b := randomMatrix(256, 256)
		bench.StartTimer()
		MulSimple(a, b)
    }
}

func BenchmarkMulGomatrix_256(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(256, 256)
		A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
		b := randomMatrix(256, 256)
		B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
		bench.StartTimer()
		A.TimesDense(B)
    }
}

func BenchmarkMulStrassPar128(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(128, 128)
		b := randomMatrix(128, 128)
		bench.StartTimer()
		MulStrassenPar(a, b)
    }
}

func BenchmarkMulDouglas__128(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(128, 128)
		b := randomMatrix(128, 128)
		bench.StartTimer()
		MulDouglas(a, b)
    }
}

func BenchmarkMulStrassen_128(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(128, 128)
		b := randomMatrix(128, 128)
		bench.StartTimer()
		MulStrassen(a, b)
    }
}

func BenchmarkMulBLAS_____128(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(128, 128)
		b := randomMatrix(128, 128)
		bench.StartTimer()
		MulBLAS(a, b)
    }
}

func BenchmarkMulSimple___128(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(128, 128)
		b := randomMatrix(128, 128)
		bench.StartTimer()
		MulSimple(a, b)
	}
}

func BenchmarkMulGomatrix_128(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(128, 128)
		A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
		b := randomMatrix(128, 128)
		B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
		bench.StartTimer()
		A.TimesDense(B)
    }
}

func BenchmarkMulBLAS______64(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(64, 64)
		b := randomMatrix(64, 64)
		bench.StartTimer()
		MulBLAS(a, b)
    }
}

func BenchmarkMulSimple____64(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(64, 64)
		b := randomMatrix(64, 64)
		bench.StartTimer()
		MulSimple(a, b)
	}
}

func BenchmarkMulGomatrix__64(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(64, 64)
		A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
		b := randomMatrix(64, 64)
		B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
		bench.StartTimer()
		A.TimesDense(B)
    }
}

func BenchmarkMulBLAS______32(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(32, 32)
		b := randomMatrix(32, 32)
		bench.StartTimer()
		MulBLAS(a, b)
    }
}

func BenchmarkMulSimple____32(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(32, 32)
		b := randomMatrix(32, 32)
		bench.StartTimer()
		MulSimple(a, b)
	}
}

func BenchmarkMulGomatrix__32(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bench.StopTimer()
		rand.Seed(time.Now().Unix())
		a := randomMatrix(32, 32)
		A := gomatrix.MakeDenseMatrix(a.data, a.height, a.width)
		b := randomMatrix(32, 32)
		B := gomatrix.MakeDenseMatrix(b.data, b.height, b.width)
		bench.StartTimer()
		A.TimesDense(B)
    }
}
