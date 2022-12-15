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

	// matrix, _ := NewMatrix(
	// 	Row[int]{1, 0, -3},
	// 	Row[int]{-2, 4, 1},
	// 	Row[int]{-5, 0, 3},
	// )
	// matrix2, _ := NewMatrix(
	// 	Row[int]{2, -1, 0},
	// 	Row[int]{3, 0, 0},
	// 	Row[int]{-5, 2, 0},
	// )
	// identity, _ := NewMatrix(
	// 	Row[int]{1, 0, 0},
	// 	Row[int]{0, 1, 0},
	// 	Row[int]{0, 0, 1},
	// )
	// Multiply(IdentityMatrix[int](3), matrix).Print()
	// matrix.ConditionalSet(0, func(val uint) bool { return val >= 30 }).Print()
	IdentityMatrix[int](5).Print()
}
