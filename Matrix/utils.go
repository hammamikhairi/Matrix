package Matrix

import (
	"fmt"
	"reflect"
)

type Value interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128
}

type Row[T Value] []T
type Diag[T Value] Row[T]
type ComplexFunction[T Value] func(value T) T
type ConditionFunction[T Value] func(value T) bool

const (
	AdditionError       = "Incompatible dimentions %d/%d and %d/%d for addition."
	MultiplicationError = "Cant multiply %d/%d by %d/%d; %d and %d must match."
)

func getIndex(i, j, rows int) int {
	return i*rows + j
}

func getPosition(index, rows, cols int) []int {
	return []int{int(index / rows), int(index % cols)}
}

func checkDimentions[T Value](M, N *Matrix[T]) {
	if N.cols != M.cols || N.rows != M.rows {
		panic(fmt.Sprintf(AdditionError, N.rows, N.cols, M.rows, M.cols))
	}
}

func genericTypeAssert[T Value, Asserted Value]() bool {
	return reflect.TypeOf(new(T)) == reflect.TypeOf(new(Asserted))
}

func newEmptyMatrix[T Value](rows, cols int) (matrix *Matrix[T]) {
	matrix = new(Matrix[T])
	matrix.cols = cols
	matrix.rows = rows
	matrix.values = make([]T, cols*rows)
	return
}

func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

func IdentityMatrix[T Value](rows int) (identity *Matrix[T]) {

	identity = newEmptyMatrix[T](rows, rows)
	for index := 0; index < rows*rows; index += rows + 1 {

		identity.values[index] = 1
	}

	return
}
