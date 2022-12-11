package main

import (
	. "github.com/hammamikhairi/Matrix/Matrix"
)

func main() {

	matrix, err := NewMatrix(
		Row[complex128]{2, 7i, 5 + 8i},
		Row[complex128]{20, 7i, 9},
		Row[complex128]{17i, 7i, 16i},
	)
	if err != nil {
		panic(err)
	}

	matrix.Print()

	matrix.Multiply(7i + 9)
	matrix.Print()

	matrix.Transpose().Print()

}
