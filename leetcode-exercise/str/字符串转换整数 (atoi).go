package str

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	return convert(clean(s))
}

func convert(sign int, abs string) int {
	num := 0
	for _, ch := range abs {
		num = num * 10 + int(ch - '0')
		if num * sign < math.MinInt32 {
			return  math.MinInt32
		}else if num * sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return num * sign
}

func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return
	}

	switch s[0] {
	case '0','1','2','3','4','5','6','7','8','9':
		sign = 1
		abs = s
	case '+':
		sign = 1
		abs = s[1:]
	case '-':
		sign = -1
		abs = s[1:]
	default:
		abs = ""
		return
	}

	for idx,ch := range abs {
		if ch < '0' || ch > '9' {
			abs = abs[:idx]
			return
		}
	}
	return
}

