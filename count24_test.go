package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	num := []float64{3, 3, 8, 8}
	str := []string{"3", "3", "8", "8"}

	if !dsf2(num, str, false, 24.0) {
		t.Error("failed@", num)
	}
}
