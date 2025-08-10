package sort

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	// 排序 按左边界升序, 右边界降序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ptr := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[ptr][0] && intervals[i][1] <= intervals[ptr][1] {
			// 可以合并
			intervals[ptr][1] = max(intervals[i][1], intervals[ptr][1])
		} else {
			ptr++
			intervals[ptr] = intervals[i]
		}
	}
	return ptr + 1
}

// https://leetcode.cn/problems/remove-covered-intervals/

// {1, 2}, {1, 4}, {3, 4}
func Show_removeCoveredIntervals() {
	intervals := [][]int{{34335, 39239}, {15875, 91969}, {29673, 66453}, {53548, 69161}, {40618, 93111}}
	res := removeCoveredIntervals(intervals)
	println(res)
}
