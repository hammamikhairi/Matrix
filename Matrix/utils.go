package Matrix

import (
	"reflect"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Value interface {
	Number |
		complex64 | complex128
}

type Row[T Value] []T
type Diag[T Value] Row[T]
type SomeFunction[T Value] func(value T) T
type ConditionFunction[T Value] func(value T) bool

const (
	AdditionError       = "Incompatible dimentions %d/%d and %d/%d for addition."
	MultiplicationError = "Cant multiply %d/%d by %d/%d; %d and %d must match."
	ReshapeError        = "Cant reshape a matrix into that form make sure the number of values is the same."
)

func getIndex(i, j, rows int) int {
	return i*rows + j
}

func getPosition(index, rows, cols int) []int {
	return []int{int(index / (rows)), int(index % cols)}
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

func toInt[T Value](number T) int64 {
	return reflect.ValueOf(number).Int()
}
func toFloat[T Value](number T) float64 {
	return reflect.ValueOf(number).Float()
}
func toUInt[T Value](number T) uint64 {
	return reflect.ValueOf(number).Uint()
}

func IsBigger[T Value](first, second T) bool {

	switch reflect.TypeOf(first).Kind() {
	case reflect.Int8, reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return toInt(first) > toInt(second)
	case reflect.Float32, reflect.Float64:
		return toFloat(first) > toFloat(second)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return toUInt(first) > toUInt(second)
	default:
		panic("not yet dummy")
	}

}

func Min[T Value](args ...T) (min T) {
	min = args[0]
	for _, val := range args {
		if IsBigger(min, val) && min != val {
			min = val
		}
	}
	return
}

func Max[T Value](args ...T) (max T) {
	max = args[0]
	for _, val := range args {
		if !IsBigger(max, val) && max != val {
			max = val
		}
	}
	return
}

func MinArg[T Value](args ...T) (ind int) {
	min := args[0]
	ind = 0
	for index, val := range args {
		if IsBigger(min, val) && min != val {
			min = val
			ind = index
		}
	}
	return
}

func MaxArg[T Value](args ...T) (ind int) {
	max := args[0]
	ind = 0
	for index, val := range args {
		if !IsBigger(max, val) && max != val {
			max = val
			ind = index
		}
	}
	return
}
