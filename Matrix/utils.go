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
type ComplexFunction[T Value] func(value complex128) T

const (
	AdditionError = "Incompatible dimentions %d/%d and %d/%d"
)

func checkDimentions[T Value](M, N *Matrix[T]) {
	if N.cols != M.cols || N.rows != M.rows {
		panic(fmt.Sprintf(AdditionError, N.rows, N.cols, M.rows, M.cols))
	}
}

func genericTypeAssert[T Value, Asserted Value]() bool {
	return reflect.TypeOf(new(T)) == reflect.TypeOf(new(Asserted))
}
