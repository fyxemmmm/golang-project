package bitOperation

func missingNumber(nums []int) int {
	res := 0

	for idx, num := range nums {
		res ^= idx ^ num
	}

	res ^= len(nums)
	return res
}
