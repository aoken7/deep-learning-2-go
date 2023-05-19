package main

import (
	"fmt"
	"numgo/numgo"
)

type Layer interface {
	forward(numgo.Mx) numgo.Mx
	getParams() []numgo.Mx
}

type Sigmoid struct {
	params []numgo.Mx
}

func (s *Sigmoid) forward(x numgo.Mx) numgo.Mx {
	return x.Sigmoid()
}

func (s *Sigmoid) getParams() []numgo.Mx {
	return s.params
}

type Affine struct {
	params []numgo.Mx
}

func (a *Affine) forward(x numgo.Mx) numgo.Mx {
	w := a.params[0]
	b := a.params[1]
	out := numgo.Add(numgo.Dot(x, w), b)
	return out
}

func (a *Affine) getParams() []numgo.Mx {
	return a.params
}

type twoLayerNet struct {
	i, h, o int

	w1 numgo.Mx
	b1 numgo.Mx
	w2 numgo.Mx
	b2 numgo.Mx

	layers []Layer
	params []numgo.Mx
}

func NewTwoLayerNet(inputSize, hiddenSize, outPutSize int) *twoLayerNet {
	i, h, o := inputSize, hiddenSize, outPutSize

	t := new(twoLayerNet)
	t.w1 = numgo.Randn(i, h)
	t.b1 = numgo.Randn(h, 1)
	t.w2 = numgo.Randn(h, o)
	t.b2 = numgo.Randn(o, 1)

	t.layers = []Layer{
		&Affine{params: []numgo.Mx{t.w1, t.b1}},
		&Sigmoid{},
		&Affine{params: []numgo.Mx{t.w2, t.b2}},
	}

	for _, l := range t.layers {
		t.params = append(t.params, l.getParams()...)
	}

	return t
}

func (t *twoLayerNet) predict(x numgo.Mx) numgo.Mx {
	for _, l := range t.layers {
		x = l.forward(x)
	}
	return x
}

func main() {
	x := numgo.Randn(10, 2)
	model := NewTwoLayerNet(2, 4, 3)
	s := model.predict(x)

	for _, v := range s.Vec {
		fmt.Println(v)
	}
}
