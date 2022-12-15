package main

import (
	. "github.com/hammamikhairi/Matrix/Matrix"
)

func main() {

	matrix, err := NewMatrix(
		Row[complex128]{16i, 7i, 5 + 8i},
		Row[complex128]{20, 7i, 9},
		Row[complex128]{17i, 7i, 16i},
	)
	if err != nil {
		panic(err)
	}

	// matrix.Print()

	// matrix.Multiply(7i + 9)
	// matrix.Print()

	// matrix.Transpose().Print()

	// fmt.Println(matrix.ComplexReals())
	// fmt.Println(matrix.ComplexImags())
	// fmt.Println(matrix.ComplexApply(func(N complex128) complex128 {
	// 	fmt.Println(N)
	// 	return N
	// }))

	// mat, _ := NewMatrix(
	// 	Row[uint]{6, 36, 33},
	// )

	matrix.Is(16i).Print()

}
