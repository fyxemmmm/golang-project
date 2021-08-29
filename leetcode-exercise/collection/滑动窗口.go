package collection

import "fmt"

func hdck1() {
	k:=5
	a := make([]int, 100)

	// 4 4-5+1
	for i := k-1; i < len(a); i++ {
		fmt.Println(i, i-k+1)
		//ans = min(ans, a[i]-a[i-k+1])
	}
}

// 4 0
// 5 1
// 6 2
// ...
// 99 95
