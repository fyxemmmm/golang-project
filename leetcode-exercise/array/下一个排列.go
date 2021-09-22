package array

func nextPermutation(nums []int)  {
	i := len(nums)-2
	for i >= 0 && nums[i+1] <= nums[i] {
		i --
	}
	if i >= 0 {
		j := len(nums)-1
		for j >= 0 && nums[j] <= nums[i] {
			j --
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	
	left := i + 1
	right := len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left ++
		right --
	}
}