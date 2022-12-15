package Matrix

import "reflect"

func (M *Matrix[T]) Trace() (sum T) {
	if M.cols != M.rows {
		panic("not square!")
	}

	for i := 0; i < M.rows; i++ {
		sum += M.GetElement(i, i)
	}
	return
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

	matrix = newEmptyMatrix[float64](M.rows, M.cols)
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

func (M *Matrix[T]) Pow(pow int) (res *Matrix[T]) {

	res = Multiply(M, M)

	for i := 2; i < pow; i++ {
		res = Multiply(res, M)
	}

	return
}

func (M *Matrix[T]) Min(axis int) (res []T) {

	switch axis {
	case 0:
		return func() []T {
			for i := 0; i < M.rows*M.cols; i += M.cols {
				temp := []T{}
				for j := 0; j < M.cols; j++ {
					temp = append(temp, M.values[i+j])
				}
				res = append(res, Min(temp...))
			}
			return res
		}()
	case 1:
		return func() []T {
			for j := 0; j < M.cols; j++ {
				temp := []T{}
				for i := 0; i < M.rows*M.cols; i += M.cols {
					temp = append(temp, M.values[i+j])
				}
				res = append(res, Min(temp...))
			}
			return res
		}()
	case -1:
		return []T{Min(M.values...)}
	}

	return
}

func (M *Matrix[T]) Max(axis int) (res []T) {

	switch axis {
	case 0:
		return func() []T {
			for i := 0; i < M.rows*M.cols; i += M.cols {
				temp := []T{}
				for j := 0; j < M.cols; j++ {
					temp = append(temp, M.values[i+j])
				}
				res = append(res, Max(temp...))
			}
			return res
		}()
	case 1:
		return func() []T {
			for j := 0; j < M.cols; j++ {
				temp := []T{}
				for i := 0; i < M.rows*M.cols; i += M.cols {
					temp = append(temp, M.values[i+j])
				}
				res = append(res, Max(temp...))
			}
			return res
		}()
	case -1:
		return []T{Max(M.values...)}
	}

	return
}

func (M *Matrix[T]) ArgMin() []int {
	return getPosition(MinArg(M.values...), M.rows, M.cols)
}

func (M *Matrix[T]) ArgMax() []int {
	return getPosition(MaxArg(M.values...), M.rows, M.cols)
}

func (M *Matrix[T]) Clip(min, max, new T) *Matrix[T] {

	if IsBigger(min, max) {
		panic("cant do that")
	}

	return M.Apply(func(value T) T {
		if IsBigger(value, min) && IsBigger(max, value) {
			return value
		} else {
			return new
		}
	})

}

func (M *Matrix[T]) Cumsum() (res *Matrix[T]) {
	res = M.Copy()

	for index := 1; index < res.Size(); index++ {
		res.values[index] = res.values[index-1] + res.values[index]
	}

	return
}

func (M *Matrix[T]) Sum() (sum T) {
	sum = 0
	for _, val := range M.values {
		sum += val
	}
	return
}

func (M *Matrix[T]) Prod() (sum T) {
	sum = 1
	for _, val := range M.values {
		sum *= val
	}
	return
}

func (M *Matrix[T]) NonZero(val T) (matrix *BoolMatrix) {

	matrix = new(BoolMatrix)
	matrix.rows = M.rows
	matrix.cols = M.cols
	matrix.values = make([]bool, matrix.rows*matrix.cols)

	for index, value := range M.values {
		matrix.values[index] = value != 0
	}

	return

}
