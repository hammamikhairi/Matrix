package Matrix

import (
	"errors"
	"fmt"
	"io"
	"math"
)

type Matrix[T Value] struct {
	rows   int
	cols   int
	values []T
}

func NewMatrix[T Value](rows ...Row[T]) (matrix *Matrix[T], err error) {

	matrix = new(Matrix[T])
	matrix.rows = len(rows)
	matrix.cols = len(rows[0])
	for _, row := range rows {
		if len(row) != matrix.cols {
			return nil, errors.New("length missmatch")
		}
		matrix.values = append(matrix.values, row...)
	}
	return matrix, nil
}

func (M *Matrix[T]) RowCount() int {
	return M.rows
}

func (M *Matrix[T]) ColCount() int {
	return M.cols
}

func (M *Matrix[T]) GetElement(i, j int) T {
	return M.values[getIndex(i, j, M.cols)]
}

func (M *Matrix[T]) SetElment(i int, j int, v T) {
	M.values[getIndex(i, j, M.cols)] = v
}

func (M *Matrix[T]) Print() {
	for i := 0; i < M.rows; i++ {
		for j := 0; j < M.cols; j++ {
			fmt.Print(M.GetElement(i, j))
			print(" ")
		}
		print("\n")
	}
}

func (M *Matrix[T]) Copy() (matrix *Matrix[T]) {
	matrix = newEmptyMatrix[T](M.rows, M.cols)
	for i := 0; i < len(matrix.values); i++ {
		matrix.values[i] = M.values[i]
	}
	return
}

func (M *Matrix[T]) Diagonal() (diag Diag[T]) {
	diag = make(Diag[T], int(math.Min(float64(M.cols), float64(M.rows))))
	for i := 0; i < len(diag); i++ {
		diag[i] = M.GetElement(i, i)
	}
	return
}

func (M *Matrix[T]) Multiply(scalar T) {
	for i := 0; i < len(M.values); i++ {
		M.values[i] *= scalar
	}
}

func (M *Matrix[T]) Apply(fn SomeFunction[T]) (matrix *Matrix[T]) {

	matrix = new(Matrix[T])
	matrix.cols = M.cols
	matrix.rows = M.rows
	matrix.values = make([]T, matrix.cols*matrix.rows)
	for i := 0; i < len(matrix.values); i++ {
		matrix.values[i] = fn(M.values[i])
	}

	return
}

func (M *Matrix[T]) IsFloat() bool {
	return genericTypeAssert[T, float32]() || genericTypeAssert[T, float64]()
}

func (M *Matrix[T]) IsUInt() bool {
	return genericTypeAssert[T, uint]() || genericTypeAssert[T, uint8]() || genericTypeAssert[T, uint16]() || genericTypeAssert[T, uint32]() || genericTypeAssert[T, uint64]()
}

func (M *Matrix[T]) IsInt() bool {
	return genericTypeAssert[T, int]() || genericTypeAssert[T, int8]() || genericTypeAssert[T, int16]() || genericTypeAssert[T, int32]() || genericTypeAssert[T, int64]() || M.IsUInt()
}

func (M *Matrix[T]) IsComplex() bool {
	return genericTypeAssert[T, complex128]() || genericTypeAssert[T, complex64]()
}

func (M *Matrix[T]) HasValue(val T) bool {
	for _, value := range M.values {
		if value == val {
			return true
		}
	}

	return false
}

func (M *Matrix[T]) ValueInMatrix(val T) (matrix *BoolMatrix) {

	matrix = new(BoolMatrix)
	matrix.rows = M.rows
	matrix.cols = M.cols
	matrix.values = make([]bool, matrix.rows*matrix.cols)

	for index, value := range M.values {
		matrix.values[index] = value == val
	}

	return

}

func (M *Matrix[T]) Filter(mesh *BoolMatrix) (values []T) {
	if M.cols != mesh.cols || M.rows != mesh.rows {
		panic("wtf bro")
	}

	for index := range M.values {
		if mesh.values[index] {
			values = append(values, M.values[index])
		}
	}

	return
}

func (M *Matrix[T]) FilterByCondition(condition ConditionFunction[T]) (matrix *BoolMatrix) {

	matrix = new(BoolMatrix)
	matrix.rows = M.rows
	matrix.cols = M.cols
	matrix.values = make([]bool, matrix.rows*matrix.cols)

	for i := 0; i < len(matrix.values); i++ {
		if condition(M.values[i]) {
			matrix.values[i] = true
		} else {
			matrix.values[i] = false
		}
	}

	return
}

func (M *Matrix[T]) ConditionalSet(value T, condition ConditionFunction[T]) (matrix *Matrix[T]) {

	matrix = newEmptyMatrix[T](M.rows, M.cols)
	for i := 0; i < len(matrix.values); i++ {
		if condition(M.values[i]) {
			matrix.values[i] = M.values[i]
		} else {
			matrix.values[i] = value
		}
	}

	return
}

func (M *Matrix[T]) IsEqualTo(N *Matrix[T]) bool {

	if M.cols != N.cols || M.rows != N.rows {
		return false
	}

	for index := range M.values {
		if M.values[index] != N.values[index] {
			return false
		}
	}
	return true
}

func (M *Matrix[T]) Flatten() []T {
	return M.values
}

func (M *Matrix[T]) Shape() []int {
	return []int{M.rows, M.cols}
}

func (M *Matrix[T]) Size() int {
	return M.cols * M.rows
}

func (M *Matrix[T]) Reshape(rows, cols int) (*Matrix[T], error) {
	assert(M.Size() == rows*cols, ReshapeError)

	reshaped := M.Copy()
	reshaped.rows, reshaped.cols = rows, cols

	return reshaped, nil
}

func (M *Matrix[T]) ToFile(stream io.ByteWriter) error {
	for i := 0; i < M.rows; i++ {
		for j := 0; j < M.cols; j++ {
			for _, byt := range fmt.Sprint(M.GetElement(i, j)) {
				stream.WriteByte(byte(byt))
			}
			// print(" ")
		}
		// print("\n")
	}

	return nil
}
