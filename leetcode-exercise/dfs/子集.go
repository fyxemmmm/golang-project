package dfs


func subsets(nums []int) [][]int {
	res := [][]int{}

	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)

		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return res

}