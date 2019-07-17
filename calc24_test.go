package main

import (
	"fmt"
	"testing"
)

func Test_dsf(t *testing.T) {
	num := []float64{3, 3, 8, 8}
	str := []string{"3", "3", "8", "8"}

	if !dsf2(num, str, false, 24.0) {
		t.Error("failed@", num)
	}
}

func Test_divzero(t *testing.T) {
	a := 1.0
	b := 0.0
	c := a / b
	fmt.Println("c:", c)
}
