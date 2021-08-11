package pointer


func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	maxLeft := 0
	maxRight := 0

	left := 0
	right := len(height)-1

	for left < right {
		maxLeft = max(height[left], maxLeft)
		maxRight = max(height[right], maxRight)

		if height[left] < height[right] {
			res += maxLeft - height[left]
			left ++
		}else {
			res += maxRight - height[right]
			right --
		}
	}

	return res
}
