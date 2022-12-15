package main

import (
	. "github.com/hammamikhairi/Matrix/Matrix"
)

func main() {

	// matrix, err := NewMatrix(
	// 	Row[complex128]{9, 2i, 5 + 8i},
	// 	Row[complex128]{9, 7i, 9},
	// 	Row[complex128]{17i, 99, 16i},
	// )
	// if err != nil {
	// 	panic(err)
	// }

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

	matrix, _ := NewMatrix(
		Row[uint]{6, 36, 33},
	)
	matrix.Print()
	matrix.ConditionalSet(0, func(val uint) bool { return val >= 30 }).Print()
}
