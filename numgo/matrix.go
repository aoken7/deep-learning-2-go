package numgo

import (
	"math"
	"math/rand"
	"time"
)

type Mx struct {
	Vec [][]float32
	T   [][]float32
}

func NewMx(a [][]float32) Mx {
	m := &Mx{}
	m.Vec = a

	aRowSize := len(a[0])
	aLineSize := len(a)
	b := make([][]float32, aRowSize)
	for i := 0; i < aRowSize; i++ {
		b[i] = make([]float32, aLineSize)
		for j := 0; j < aLineSize; j++ {
			b[i][j] = a[j][i]
		}
	}
	m.T = b

	return *m
}

func (m *Mx) shape() []int {
	s := make([]int, 2)
	s[0] = len(m.Vec[0]) // 行数
	s[1] = len(m.Vec)    // 列数
	return s
}

// (line, row)の大きさでmxをランダムに初期化
func Randn(l, r int) Mx {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	m := make([][]float32, l)
	for i := 0; i < l; i++ {
		m[i] = make([]float32, r)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < r; j++ {
			m[i][j] = rand.Float32() + 0.5
		}
	}

	return NewMx(m)
}

func Exp(a Mx) Mx {
	aShape := a.shape()
	b := make([][]float32, aShape[1])

	for i := 0; i < aShape[1]; i++ {
		b[i] = make([]float32, aShape[0])
		for j := 0; j < aShape[0]; j++ {
			b[i][j] = float32(math.Exp(float64(a.Vec[i][j])))
		}
	}

	return NewMx(b)
}

func AddScl(a Mx, scalar float32) Mx {
	aShape := a.shape()
	b := make([][]float32, aShape[1])

	for i := 0; i < aShape[1]; i++ {
		b[i] = make([]float32, aShape[0])
		for j := 0; j < aShape[0]; j++ {
			b[i][j] = a.Vec[i][j] + scalar
		}
	}
	return NewMx(b)
}

func Dot(a, b Mx) Mx {
	a_shape := a.shape()
	b_shape := b.shape()

	if a_shape[0] != b_shape[1] {
		panic("Matrix shape does not match.")
	}

	c := make([][]float32, a_shape[1])
	for i := 0; i < a_shape[1]; i++ {
		c[i] = make([]float32, b_shape[0])
	}

	for i := 0; i < a_shape[1]; i++ {
		for j := 0; j < b_shape[0]; j++ {
			for k := 0; k < a_shape[0]; k++ {
				c[i][j] += a.Vec[i][k] * b.Vec[k][j]
			}
		}
	}

	return NewMx(c)
}
