package str

func reverseString(s []byte)  {
	length := len(s)
	for i := 0; i < length/2; i ++ {
		s[length-i-1], s[i] = s[i], s[length-i-1]
	}
}
