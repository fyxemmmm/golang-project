package binarySearch

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left := 0
	right := x
	for left < right {
		mid := (right+left) >> 1
		if mid * mid == x {
			return mid
		}else if mid * mid < x {
			left = mid + 1
		}else {
			right = mid
		}
	}

	return left - 1
}