package dfs


func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for _, num := range nums {
			if visited[num] == true {
				continue
			}

			visited[num] = true
			path = append(path, num)

			dfs(path)

			visited[num] = false
			path = path[:len(path)-1]
		}

	}

	dfs([]int{})
	return res
}
