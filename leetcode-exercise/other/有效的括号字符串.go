package other
func checkValidString(s string) bool {
	leftCount, rightCount := 0, 0
	for i := 0; i < len(s); i ++ {
		c := s[i]
		if c == ')' {
			rightCount ++
		}else {
			leftCount ++
		}
		if leftCount < rightCount {
			return false
		}
	}

	leftCount, rightCount = 0, 0
	for i := len(s)-1; i >= 0; i -- {
		c := s[i]
		if c == '(' {
			leftCount ++
		}else {
			rightCount ++
		}

		if rightCount < leftCount {
			return false
		}
	}
	return true

}
