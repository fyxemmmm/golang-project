package other

func rotate(matrix [][]int)  {
	if len (matrix) == 0 {
		return
	}

	length := len(matrix)
	for i := 0; i < length/2; i ++ {
		matrix[length-i-1], matrix[i] = matrix[i], matrix[length-i-1]
	}

	for i := 0; i < length; i ++ {
		for j := 0; j < i; j ++ {
			matrix[i][j],matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
