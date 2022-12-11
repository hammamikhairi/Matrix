package Matrix

import "fmt"

type Value interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float64 | float32 |
	complex64 | complex128
}

type Row[T Value] []T

type Diag[T Value] Row[T]

const (
	AdditionError = "Incompatible dimentions %d/%d and %d/%d"
)

func checkDimentions[T Value](M, N *Matrix[T]) {
	if N.cols != M.cols || N.rows != M.rows {
		panic(fmt.Sprintf(AdditionError, N.rows, N.cols, M.rows, M.cols))
	}
}
