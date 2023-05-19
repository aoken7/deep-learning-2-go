package common

import "numgo/numgo"

type Layer interface {
	Forward(numgo.Mx) numgo.Mx
	GetParams() []numgo.Mx
}

type Sigmoid struct {
	Params []numgo.Mx
}

func (s *Sigmoid) Forward(x numgo.Mx) numgo.Mx {
	return x.Sigmoid()
}

func (s *Sigmoid) GetParams() []numgo.Mx {
	return s.Params
}

type Affine struct {
	Params []numgo.Mx
}

func (a *Affine) Forward(x numgo.Mx) numgo.Mx {
	w := a.Params[0]
	b := a.Params[1]
	out := numgo.Add(numgo.Dot(x, w), b)
	return out
}

func (a *Affine) GetParams() []numgo.Mx {
	return a.Params
}
