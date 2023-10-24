package leetcode_go

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	ans = append(ans, intervals[0])

	for _, interval := range intervals[1:] {
		if ans[len(ans)-1][1] < interval[0] {
			ans = append(ans, interval)
		} else {
			if ans[len(ans)-1][1] > interval[1] {
				continue
			}
			ans[len(ans)-1][1] = interval[1]
		}
	}
	return ans
}
