package array

import "sort"

func mergeArea(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	prev := intervals[0]
	for i := 1; i < len(intervals); i ++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		}else {
			prev[1] = max(cur[1], prev[1])
		}
	}
	res = append(res, prev)
	return res
}
