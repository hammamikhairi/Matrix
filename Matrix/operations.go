package Matrix

// TODO -> implement These :
//
//	flatten() -> return values slice
//	shape() -> (row, col)
//	size() -> num of values
//	real() -> the real part of a complex mat
//	im() -> the imaginary part of a complex mat
//
//	min(axis) / max(axis) -> return the min/max values can specify the axis
//	argmin() / argmax() -> return the index of the min/max values
//  clip(min, max) -> remove values not in range min <-> max
//	cumsum -> commulative sum
//	sum() / prod()
//
//	reshape(rows, cols)
//	tofile(io.writer)
//
//	nonzero() -> returns the elements that are not zero
//  inverse of the above
//
//  Is(val) -> return the indecies of the values matching <val>
//  IsNot(val) -> inverse of the above

func Add[T Value](M, N *Matrix[T]) (final *Matrix[T]) {

	checkDimentions(M, N)

	final = M.Copy()
	for i := 0; i < len(M.values); i++ {
		final.values[i] = M.values[i] + N.values[i]
	}

	return
}

func Substract[T Value](M, N *Matrix[T]) (final *Matrix[T]) {

	checkDimentions(M, N)

	final = M.Copy()
	for i := 0; i < len(M.values); i++ {
		final.values[i] = M.values[i] - N.values[i]
	}

	return
}
