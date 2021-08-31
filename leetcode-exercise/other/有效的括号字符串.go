package other

func checkValidString(s string) bool {
	left := 0
	right := len(s)-1

	if len(s) == 1 {
		if s[0] == '*'{
			return true
		}

		return false
	}

	if len(s) == 2 {
		if s[left] == '*' || s[right] == '*' {
			return true
		}
		if s[left] == '(' && s[right] == ')' {
			return true
		}

		return false
	}

	for left <= right {
		if right-left == 0 || right -left == 1{
			return true
		}

		if s[left] == '(' && s[right] == '(' {
			return false
		}

		if s[left] == ')' && s[right] == ')' {
			return false
		}

		if s[left] == ')' && s[right] == '(' {
			return false
		}


		left ++
		right --
	}


	return false

}
