package Matrix

import "fmt"

// TODO -> implement These :
//
//$	flatten() -> return values slice
//$	shape() -> (row, col)
//$	size() -> num of values
//$	ComplexReals() -> the real part of a complex mat
//$	ComplexImags() -> the imaginary part of a complex mat
//
//$	min(axis) / max(axis) -> return the min/max values can specify the axis
//$	argmin() / argmax() -> return the index of the min/max values
//$  clip(min, max) -> remove values not in range min <-> max
//$	cumsum -> commulative sum
//$	sum() / prod()
//
//$	reshape(rows, cols)
//	tofile(io.writer)
//
//$	nonzero() -> returns the elements that are not zero
//$  inverse of the above
//
//  Is(val) -> return the indecies of the values matching <val>
//  IsNot(val) -> inverse of the above

func Add[T Value](M, N *Matrix[T]) (final *Matrix[T]) {

	assert(N.cols != M.cols || N.rows != M.rows, fmt.Sprintf(AdditionError, N.rows, N.cols, M.rows, M.cols))

	final = M.Copy()
	for i := 0; i < len(M.values); i++ {
		final.values[i] = M.values[i] + N.values[i]
	}

	return
}

func Substract[T Value](M, N *Matrix[T]) (final *Matrix[T]) {

	assert(N.cols != M.cols || N.rows != M.rows, fmt.Sprintf(AdditionError, N.rows, N.cols, M.rows, M.cols))

	final = M.Copy()
	for i := 0; i < len(M.values); i++ {
		final.values[i] = M.values[i] - N.values[i]
	}

	return
}

func Multiply[T Value](M, N *Matrix[T]) (final *Matrix[T]) {

	assert(M.cols == N.rows, fmt.Sprintf(MultiplicationError, M.rows, M.cols, N.rows, N.cols, M.cols, N.rows))

	final = newEmptyMatrix[T](M.rows, N.cols)

	for i := 0; i < M.rows; i++ {
		for k := 0; k < N.cols; k++ {
			localRes := T(0)
			for j := 0; j < M.cols; j++ {
				localRes += M.GetElement(i, j) * N.GetElement(j, k)
			}
			final.values[getIndex(i, k, M.rows)] = localRes
		}
	}

	return
}
