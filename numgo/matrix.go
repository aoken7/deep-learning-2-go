package numgo

import (
	"math"
	"math/rand"
)

type Mx struct {
	Vec [][]float32
	t   [][]float32
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
	m.t = b

	return *m
}

// [列, 行]で行列の形状を返す
func (m *Mx) Shape() []int {
	s := make([]int, 2)
	s[0] = len(m.Vec)    // 行数
	s[1] = len(m.Vec[0]) // 列数
	return s
}

func (m *Mx) T() Mx {
	return NewMx(m.t)
}

func (m *Mx) Max(axis int) Mx {
	mShape := m.Shape()

	a := [][]float32{}

	// 行列の列ごとのMaxを計算
	if axis == 0 {
		a = make([][]float32, 1)
		a[0] = make([]float32, mShape[1])
		// Maxをとるのでスライスを-infで初期化
		for k := 0; k < mShape[1]; k++ {
			a[0][k] = float32(math.Inf(-1))
		}

		for i := 0; i < mShape[0]; i++ {
			for j := 0; j < mShape[1]; j++ {
				if a[0][j] < m.Vec[i][j] {
					a[0][j] = m.Vec[i][j]
				}
			}
		}
		// 行列の行ごとのMaxを計算
	} else if axis == 1 {
		a = make([][]float32, mShape[0])
		for i := 0; i < mShape[0]; i++ {
			a[i] = make([]float32, 1)
			// Maxをとるのでスライスを-infで初期化
			a[i][0] = float32(math.Inf(-1))

			for j := 0; j < mShape[1]; j++ {
				if a[i][0] < m.Vec[i][j] {
					a[i][0] = m.Vec[i][j]
				}
			}
		}
	}

	return NewMx(a)
}

func (m *Mx) Sum(axis int) Mx {
	mShape := m.Shape()

	a := [][]float32{}

	// 行列の列ごとのSumを計算
	if axis == 0 {
		a = make([][]float32, 1)
		a[0] = make([]float32, mShape[1])

		for i := 0; i < mShape[0]; i++ {
			for j := 0; j < mShape[1]; j++ {
				a[0][j] += m.Vec[i][j]

			}
		}
		// 行列の行ごとのSumを計算
	} else if axis == 1 {
		a = make([][]float32, mShape[0])
		for i := 0; i < mShape[0]; i++ {
			a[i] = make([]float32, 1)
			for j := 0; j < mShape[1]; j++ {
				a[i][0] += m.Vec[i][j]
			}
		}
	}

	return NewMx(a)
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

// (line, row) > 0 の大きさでmxをゼロ初期化
func ZeroLike(l, r int) Mx {
	m := make([][]float32, l)
	for i := 0; i < l; i++ {
		m[i] = make([]float32, r)
	}
	return NewMx(m)
}

// 行列と列が1つの行列もしくはスカラの演算を行う
func vectorOperation(a, w Mx, fn func(float32, float32) float32) Mx {
	aShape := a.Shape()
	b := make([][]float32, aShape[0])

	for i := 0; i < aShape[0]; i++ {
		b[i] = make([]float32, aShape[1])

		for j := 0; j < aShape[1]; j++ {
			if len(w.Vec[0]) == 1 {
				//b[i][j] = a.Vec[i][j] + w.Vec[0][0]
				b[i][j] = fn(a.Vec[i][j], w.Vec[0][0])
			} else {
				//b[i][j] = a.Vec[i][j] + w.Vec[0][j]
				b[i][j] = fn(a.Vec[i][j], w.Vec[0][j])
			}
		}
	}
	return NewMx(b)
}

// wの大きさが(1x1)の時は行列とスカラの和になる
func Add(a, w Mx) Mx {
	return vectorOperation(a, w, func(f1, f2 float32) float32 {
		return f1 + f2
	})
}

func Sub(a, w Mx) Mx {
	return vectorOperation(a, w, func(f1, f2 float32) float32 {
		return f1 - f2
	})
}

func Div(a, w Mx) Mx {
	return vectorOperation(a, w, func(f1, f2 float32) float32 {
		return f1 / f2
	})
}

// 行列積を計算する。aの行とbの列の大きさを合わせる必要あり。
func Dot(a, b Mx) Mx {
	a_shape := a.Shape()
	b_shape := b.Shape()

	if a_shape[1] != b_shape[0] {
		panic("Matrix shape does not match.")
	}

	c := make([][]float32, a_shape[0])
	for i := 0; i < a_shape[0]; i++ {
		c[i] = make([]float32, b_shape[1])
		for j := 0; j < b_shape[1]; j++ {
			for k := 0; k < a_shape[1]; k++ {
				c[i][j] += a.Vec[i][k] * b.Vec[k][j]
			}
		}
	}

	return NewMx(c)
}

// 行列のアダマール積
func Had(a, b Mx) Mx {
	aShape := a.Shape()
	bShape := b.Shape()

	if aShape[0] != bShape[0] || aShape[1] != bShape[1] {
		panic("Matrix shape does not mathc.")
	}

	c := make([][]float32, aShape[0])
	for i := 0; i < aShape[0]; i++ {
		c[i] = make([]float32, aShape[1])
		for j := 0; j < aShape[1]; j++ {
			c[i][j] = a.Vec[i][j] * b.Vec[i][j]
		}
	}

	return NewMx(c)
}

// 行列の和をとって2->1次元にする
func Sum(a Mx, axis int) Mx {
	aShape := a.Shape()

	if axis < 0 || 1 < axis {
		panic("'axis' must range from 0 to 1.")
	}

	b := make([][]float32, 1)

	if axis == 0 {
		b[0] = make([]float32, aShape[1])
		for i := 0; i < aShape[0]; i++ {
			for j := 0; j < aShape[1]; j++ {
				b[0][j] += a.Vec[i][j]
			}
		}
	} else {
		b[0] = make([]float32, aShape[0])
		for i := 0; i < aShape[0]; i++ {
			for j := 0; j < aShape[1]; j++ {
				b[0][i] += a.Vec[i][j]
			}
		}
	}

	return NewMx(b)
}
