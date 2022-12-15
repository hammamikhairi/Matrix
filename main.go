package main

import (
	"fmt"

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
	// 	Row[int]{-2, -50, 1},
	// 	Row[int]{-5, 0, -700},
	// )
	// matrix2, _ := NewMatrix(
	// 	Row[int]{2, -1, 0},
	// 	Row[int]{3, 0, 0},
	// 	Row[int]{-5, 2, 0},
	// )
	matrix, _ := NewMatrix(
		Row[int]{1, 2, 3},
		Row[int]{4, 5, 6},
		Row[int]{0, 0, 1},
		Row[int]{0, 0, 1},
	)

	fmt.Println(matrix.Filter(matrix.FilterByCondition(func(value int) bool {
		return value < 1
	})))

	// fmt.Println(Multiply(matrix, matrix).IsEqualTo(matrix.Pow(3)))
	// print("\n")
	// matrix.Mins(0)
	// matrix.ConditionalSet(0, func(val uint) bool { return val >= 30 }).Print()
	// IdentityMatrix[int](5).Print()
	matrix.Print()
	// fmt.Println(matrix.Max(-1))
	// fmt.Println(getPosition(5, 2, 3))
	// matrix.Cumsum().Print()
	// fmt.Println(reflect.ValueOf(0).Int() < reflect.ValueOf(1).Int())

	// file, _ := os.Open("goal.txt")

	// her := bufio.NewWriter(file)

	// n, err := her.WriteString("here")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)

}
