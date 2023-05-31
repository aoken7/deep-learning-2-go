package main

import (
	"fmt"
	"numgo/numgo"
	"testing"
)

func printMx(m numgo.Mx) {
	for _, v := range m.Vec {
		fmt.Println(v)
	}
}

func TestMxT(t *testing.T) {
	a := numgo.NewMx([][]float32{{1, 2}, {3, 4}, {5, 6}})
	fmt.Println("origin:", a.Vec)
	fmt.Println("T:", a.T())
}

func TestDot(t *testing.T) {
	a := numgo.NewMx([][]float32{{1, 2, 3}})
	b := numgo.NewMx([][]float32{{4}, {5}, {6}})

	printMx(numgo.Dot(a, b))

	c := numgo.NewMx([][]float32{{1, 2}, {3, 4}})
	d := numgo.NewMx([][]float32{{5, 6}, {7, 8}})

	printMx(numgo.Dot(c, d))
}

func TestWeghtCalc(t *testing.T) {
	w1 := numgo.Randn(2, 4)
	b1 := numgo.Randn(1, 4)
	x := numgo.Randn(10, 2)

	h := numgo.Add(numgo.Dot(x, w1), b1)

	printMx(h)
	fmt.Println(h.Shape())
}

func TestSigmoid(t *testing.T) {
	x := numgo.Randn(10, 2)
	w1 := numgo.Randn(2, 4)
	b1 := numgo.Randn(4, 1)
	w2 := numgo.Randn(4, 3)
	b2 := numgo.Randn(3, 1)

	x = numgo.Dot(x, w1)
	fmt.Println("xShape:", x.Shape())
	fmt.Println("b1Shape:", b1.Shape())
	h := numgo.Add(x.T(), b1)
	h = h.T()
	a := h.Sigmoid()
	a = numgo.Dot(a, w2)
	s := numgo.Add(a.T(), b2)

	printMx(s)
}
