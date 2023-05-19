package main

import (
	"fmt"
	"numgo/numgo"
	"testing"
)

func TestMxT(t *testing.T) {
	a := numgo.NewMx([][]float32{{1, 2}, {3, 4}, {5, 6}})
	fmt.Println("origin:", a.Vec)
	fmt.Println("T:", a.T)
}

func TestDot(t *testing.T) {
	a := numgo.NewMx([][]float32{{1, 2, 3}})
	b := numgo.NewMx([][]float32{{4}, {5}, {6}})

	fmt.Println(numgo.Dot(a, b).Vec)

	c := numgo.NewMx([][]float32{{1, 2}, {3, 4}})
	d := numgo.NewMx([][]float32{{5, 6}, {7, 8}})

	fmt.Println(numgo.Dot(c, d).Vec)
}
