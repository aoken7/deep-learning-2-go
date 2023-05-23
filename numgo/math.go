package numgo

import "math"

func applyMx(m *Mx, fn func(float32) float32) Mx {
	aShape := m.Shape()
	b := make([][]float32, aShape[0])

	for i := 0; i < aShape[0]; i++ {
		b[i] = make([]float32, aShape[1])
		for j := 0; j < aShape[1]; j++ {
			b[i][j] = fn(m.Vec[i][j])
		}
	}

	return NewMx(b)
}

func (m *Mx) Exp() Mx {
	return applyMx(m, func(f float32) float32 {
		return float32(math.Exp(float64(f)))
	})
}

func (m *Mx) FlipSign() Mx {
	return applyMx(m, func(f float32) float32 {
		return -f
	})
}

func (m *Mx) Sigmoid() Mx {
	return applyMx(m, func(f float32) float32 {
		return 1 / (1 + float32(math.Exp(float64(-f))))
	})
}

// 行列の一要素に処理 fn を適応させる。applyMxのラップ関数
func (m *Mx) Apply(fn func(f float32) float32) Mx {
	return applyMx(m, fn)
}
