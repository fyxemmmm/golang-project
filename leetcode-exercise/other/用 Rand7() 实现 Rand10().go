package other

import "math/rand"

func rand10() int {
	for {
		num := (rand7()-1) * 7 + rand7()
		if num <= 40 {
			return num % 10 + 1
		}
	}
}




func rand7() int {
	return rand.Int()
}

