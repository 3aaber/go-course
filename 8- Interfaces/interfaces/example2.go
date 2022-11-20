package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func example2() {
	var i I

	i = &T{"Hello"}
	describe2(i)
	i.M()

	i = F(math.Pi)
	describe2(i)
	i.M()
}

func describe2(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
