package Matrix

import "fmt"

type BoolMatrix struct {
	rows   int
	cols   int
	values []bool
}

func (M *BoolMatrix) GetElement(i, j int) bool {
	return M.values[getIndex(i, j, M.cols)]
}

func (M *BoolMatrix) Print() {
	for i := 0; i < M.rows; i++ {
		for j := 0; j < M.cols; j++ {
			fmt.Print(M.GetElement(i, j))
			print(" ")
		}
		print("\n")
	}
}
