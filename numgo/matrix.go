package numgo

import (
	"math"
	"math/rand"
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

func (m *Mx) Shape() []int {
	s := make([]int, 2)
	s[0] = len(m.Vec[0]) // 行数
	s[1] = len(m.Vec)    // 列数
	return s
}

func boxMuller() float64 {
	u := rand.Float64()
	v := rand.Float64()
	z := math.Sqrt(-2*math.Log(u)) * math.Cos(2*math.Pi*v)
	return z
}

// (line, row) > 0 の大きさでmxをランダムに初期化
func Randn(l, r int) Mx {
	m := make([][]float32, l)
	for i := 0; i < l; i++ {
		m[i] = make([]float32, r)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < r; j++ {
			m[i][j] = float32(boxMuller())
		}
	}

	return NewMx(m)
}

// wの大きさが(1x1)の時は行列とスカラの和になる
func Add(a Mx, w Mx) Mx {
	aShape := a.Shape()
	b := make([][]float32, aShape[1])

	for i := 0; i < aShape[1]; i++ {
		b[i] = make([]float32, aShape[0])

		for j := 0; j < aShape[0]; j++ {
			if len(w.Vec[0]) == 1 {
				b[i][j] = a.Vec[i][j] + w.Vec[0][0]
			} else {
				b[i][j] = a.Vec[i][j] + w.Vec[0][j]
			}
		}
	}

	return NewMx(b)
}

// 行列積を計算する。aの行とbの列の大きさを合わせる必要あり。
func Dot(a, b Mx) Mx {
	a_shape := a.Shape()
	b_shape := b.Shape()

	if a_shape[0] != b_shape[1] {
		panic("Matrix shape does not match.")
	}

	c := make([][]float32, a_shape[1])
	for i := 0; i < a_shape[1]; i++ {
		c[i] = make([]float32, b_shape[0])
		for j := 0; j < b_shape[0]; j++ {
			for k := 0; k < a_shape[0]; k++ {
				c[i][j] += a.Vec[i][k] * b.Vec[k][j]
			}
		}
	}

	return NewMx(c)
}
