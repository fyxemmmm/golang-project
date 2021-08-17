package array

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}

	for idx, num := range nums {
		if index, ok := hash[target-num]; ok {
			return []int{idx, index}
		}
		hash[num] = idx
	}

	return nil
}

