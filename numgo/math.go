package numgo

import "math"

func applyMx(m *Mx, fn func(float32) float32) Mx {
	aShape := m.Shape()
	b := make([][]float32, aShape[1])

	for i := 0; i < aShape[1]; i++ {
		b[i] = make([]float32, aShape[0])
		for j := 0; j < aShape[0]; j++ {
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
