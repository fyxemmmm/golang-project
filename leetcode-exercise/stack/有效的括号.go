package stack


func isValid(s string) bool {
	stack := []byte{}
	hash := map[byte]byte{')':'(', ']': '[', '}': '{'}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, byte(ch))
		}else if len(stack) > 0 && stack[len(stack)-1] == hash[byte(ch)]{
			stack = stack[:len(stack)-1]
		}else {
			return false
		}

	}

	return len(stack) == 0
}