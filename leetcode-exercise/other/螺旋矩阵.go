package other

func spiralOrder(matrix [][]int) []int {
	res := []int{}

	top, left := 0, 0
	right, bottom := len(matrix[0])-1, len(matrix)-1
	numEle := len(matrix) * len(matrix[0])

	for numEle > 0 {
		for i := left; i <= right && numEle > 0; i ++ {
			res = append(res, matrix[top][i])
			numEle --
		}
		top ++

		for i := top; i <= bottom && numEle > 0; i ++ {
			res =append(res, matrix[i][right])
			numEle --
		}
		right --

		for i := right; i >= left && numEle > 0; i -- {
			res = append(res, matrix[bottom][i])
			numEle --
		}
		bottom --


		for i := bottom; i >= top && numEle > 0; i -- {
			res = append(res, matrix[i][left])
			numEle --

		}
		left ++

	}

	return res
}
