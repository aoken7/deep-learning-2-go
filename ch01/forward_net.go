package main

import (
	"fmt"
	"numgo/common"
	"numgo/numgo"
)

type twoLayerNet struct {
	i, h, o int

	w1 numgo.Mx
	b1 numgo.Mx
	w2 numgo.Mx
	b2 numgo.Mx

	layers []common.Layer
	params []numgo.Mx
}

func NewTwoLayerNet(inputSize, hiddenSize, outPutSize int) *twoLayerNet {
	i, h, o := inputSize, hiddenSize, outPutSize

	t := new(twoLayerNet)
	t.w1 = numgo.Randn(i, h)
	t.b1 = numgo.Randn(h, 1)
	t.w2 = numgo.Randn(h, o)
	t.b2 = numgo.Randn(o, 1)

	t.layers = []common.Layer{
		&common.Affine{Params: []numgo.Mx{t.w1, t.b1}},
		&common.Sigmoid{},
		&common.Affine{Params: []numgo.Mx{t.w2, t.b2}},
	}

	for _, l := range t.layers {
		t.params = append(t.params, l.GetParams()...)
	}

	return t
}

func (t *twoLayerNet) predict(x numgo.Mx) numgo.Mx {
	for _, l := range t.layers {
		x = l.Forward(x)
	}
	return x
}

func main() {
	x := numgo.Randn(10, 2)
	model := NewTwoLayerNet(2, 4, 3)
	s := model.predict(x)

	for _, v := range model.params {
		fmt.Println(v)
	}

	for _, v := range s.Vec {
		fmt.Println(v)
	}
}
