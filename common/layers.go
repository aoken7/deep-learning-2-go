package common

import (
	"fmt"
	"numgo/numgo"
)

type Layer interface {
	Forward(numgo.Mx) numgo.Mx
	Backward(numgo.Mx) numgo.Mx
	GetParams() []numgo.Mx
}

type Sigmoid struct {
	Params []numgo.Mx
	grads  []numgo.Mx
	out    numgo.Mx
}

func (s *Sigmoid) Forward(x numgo.Mx) numgo.Mx {
	out := x.Sigmoid()
	s.out = out
	return out
}

func (s *Sigmoid) Backward(dout numgo.Mx) numgo.Mx {
	out := s.out.Apply(func(f float32) float32 { return 1.0 - f })
	dx := numgo.Had(numgo.Had(dout, out), s.out) // dout * (1.0 - s.out) * s.out
	return dx
}

func (s *Sigmoid) GetParams() []numgo.Mx {
	return s.Params
}

type Affine struct {
	Params []numgo.Mx
	grads  []numgo.Mx
	x      numgo.Mx
}

func NewAffine(w, b numgo.Mx) Affine {
	wShape := w.Shape()
	bShape := b.Shape()

	af := new(Affine)
	af.Params = []numgo.Mx{w, b}
	af.grads = []numgo.Mx{
		numgo.ZeroLike(wShape[0], wShape[1]),
		numgo.ZeroLike(bShape[0], bShape[1]),
	}

	return *af
}

func (a *Affine) Forward(x numgo.Mx) numgo.Mx {
	w := a.Params[0]
	b := a.Params[1]
	out := numgo.Add(numgo.Dot(x, w), b)
	a.x = x
	return out
}

func (a *Affine) Backward(dout numgo.Mx) numgo.Mx {
	w := a.Params[0]
	dx := numgo.Dot(dout, w.T())
	dw := numgo.Dot(a.x.T(), dout)
	db := numgo.Sum(dout, 0)

	a.grads[0] = dw
	a.grads[1] = db

	return dx
}

func (a *Affine) GetParams() []numgo.Mx {
	return a.Params
}

type MatMul struct {
	Params []numgo.Mx
	grads  []numgo.Mx
	x      numgo.Mx
}

func NewMatMul(w numgo.Mx) MatMul {
	wShape := w.Shape()

	ma := new(MatMul)
	ma.Params = []numgo.Mx{w}
	ma.grads = []numgo.Mx{numgo.ZeroLike(wShape[0], wShape[1])}

	return *ma
}

func (m *MatMul) Forward(x numgo.Mx) numgo.Mx {
	w := m.Params[0]
	out := numgo.Dot(x, w)
	m.x = x
	return out
}

func (m *MatMul) Backward(dout numgo.Mx) numgo.Mx {
	w := m.Params[0]
	dx := numgo.Dot(dout, w.T())
	dw := numgo.Dot(m.x.T(), dout)

	m.grads[0] = dw
	return dx
}

func (m *MatMul) GetParams() []numgo.Mx {
	return m.Params
}

type SoftMaxWithLoss struct {
	params []numgo.Mx
	grads  []numgo.Mx
	out    numgo.Mx
}

func (s *SoftMaxWithLoss) Forward(x numgo.Mx) numgo.Mx {
	return x
}

func softmax(x numgo.Mx) numgo.Mx {
	fmt.Println("xShape:", x.Shape())
	tmp := x.Max(1)
	fmt.Println("x.MaxShape:", tmp.Shape())
	x = numgo.Sub(x, x.Max(1))
	x = x.Exp()
	x = numgo.Div(x, x.Sum(1))
	return x
}
