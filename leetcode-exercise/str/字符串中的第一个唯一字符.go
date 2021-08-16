package str

func firstUniqChar(s string) int {
	hash := map[int32]int{}

	for idx, ch := range s {
		hash[ch] = idx
	}

	for idx, ch := range s {
		if hash[ch] == idx {
			return idx
		}else {
			hash[ch] = -1
		}
	}

	return -1
}
