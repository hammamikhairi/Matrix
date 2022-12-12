package Matrix

import (
	"errors"
	"fmt"
	"math"
	"reflect"
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

func getIndex(i, j, rows int) int {
	return i*rows + j
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
	matrix = new(Matrix[T])
	matrix.cols = M.cols
	matrix.rows = M.rows
	matrix.values = make([]T, matrix.cols*matrix.rows)
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

func (M *Matrix[T]) Trace() (sum T) {
	if M.cols != M.rows {
		panic("not square!")
	}

	for i := 0; i < M.rows; i++ {
		sum += M.GetElement(i, i)
	}
	return
}

func (M *Matrix[T]) Multiply(scalar T) {
	for i := 0; i < len(M.values); i++ {
		M.values[i] *= scalar
	}
}

func (M *Matrix[T]) Transpose() (final *Matrix[T]) {
	final = M.Copy()
	final.cols, final.rows = M.rows, M.cols
	for i := 0; i < final.rows; i++ {
		for j := 0; j < final.cols; j++ {
			final.SetElment(i, j, M.GetElement(j, i))
		}
	}
	return
}

func (M *Matrix[T]) ComplexReals() (matrix *Matrix[float64]) {

	if !(genericTypeAssert[T, complex128]() || genericTypeAssert[T, complex64]()) {
		panic("must be complex")
	}

	matrix = new(Matrix[float64])
	matrix.cols = M.cols
	matrix.rows = M.rows
	matrix.values = make([]float64, matrix.cols*matrix.rows)
	for i := 0; i < len(matrix.values); i++ {
		matrix.values[i] = real(reflect.ValueOf(M.values[i]).Complex())
	}

	return
}

func (M *Matrix[T]) ComplexImags() (matrix *Matrix[float64]) {

	if !(genericTypeAssert[T, complex128]() || genericTypeAssert[T, complex64]()) {
		panic("must be complex")
	}

	matrix = new(Matrix[float64])
	matrix.cols = M.cols
	matrix.rows = M.rows
	matrix.values = make([]float64, matrix.cols*matrix.rows)
	for i := 0; i < len(matrix.values); i++ {
		matrix.values[i] = imag(reflect.ValueOf(M.values[i]).Complex())
	}

	return
}

func (M *Matrix[T]) ComplexApply(fn ComplexFunction[T]) (matrix *Matrix[T]) {

	if !(genericTypeAssert[T, complex128]() || genericTypeAssert[T, complex64]()) {
		panic("must be complex")
	}

	matrix = new(Matrix[T])
	matrix.cols = M.cols
	matrix.rows = M.rows
	matrix.values = make([]T, matrix.cols*matrix.rows)
	for i := 0; i < len(matrix.values); i++ {
		matrix.values[i] = fn(reflect.ValueOf(M.values[i]).Complex())
	}

	return
}
