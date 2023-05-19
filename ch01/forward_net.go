package main

import (
	"fmt"
	"numgo/numgo"
)

type Sigmoid struct {
	params interface{}
}

func (s *Sigmoid) foward(a numgo.Mx) numgo.Mx {
	b := a.Exp()
	b = numgo.Add(b, numgo.NewMx([][]float32{{1}}))
	return b
}

func main() {
	x := numgo.Randn(10, 2)
	s := &Sigmoid{}

	fmt.Println(s.foward(x))

	//fmt.Println(x)

}
