package pointer

func maxArea(height []int) int {
	left := 0
	right := len(height)-1

	res := 0
	for left < right {
		area := (right-left) *min(height[left], height[right])
		if area > res {
			res = area
		}

		if height[left] < height[right] {
			left ++
		}else {
			right --
		}

	}
	return res
}

