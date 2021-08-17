package other

func majorityElement(nums []int) int {
	res := 0
	vote := 0

	for _, num := range nums {
		if vote == 0 {
			res = num
		}

		if num == res {
			vote ++
		}else {
			vote --
		}
	}

	return res
}
