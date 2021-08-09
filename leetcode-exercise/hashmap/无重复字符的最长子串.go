package hashmap

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0
	for len(s) > 0 {
		m := map[int32]bool{}
		for _, ch := range s {
			if !m[ch] {
				m[ch] = true
				if len(m) > res {
					res = len(m)
				}
			}else {
				if len(m) > res {
					res = len(m)
				}
				break
			}
		}
		s = s[1:]
	}

	return res
}
