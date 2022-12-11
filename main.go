package main

import (
	. "github.com/hammamikhairi/Matrix/Matrix"
)

func main() {
	matrix, err := NewMatrix(
		Row[float64]{1, 2, 3},
		Row[float64]{4, 5, 6},
		Row[float64]{7, 8, 9},
	)
	if err != nil {
		panic(err)
	}

	second := matrix.Copy()
	second.Multiply(2)
	matrix.Print()
	print("\n")
	second.Print()
}
