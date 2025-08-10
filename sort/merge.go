package sort

import (
	"fmt"
	"sort"
)

func mergr(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 {
			res = append(res, intervals[i])
			continue
		}
		last := res[len(res)-1]
		cur := intervals[i]
		if last[1] >= cur[0] {
			last[1] = maxInt(cur[1], last[1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// https://leetcode.cn/problems/merge-intervals/description/
func Show_merge() {
	arr := [][]int{{1, 4}, {0, 4}}
	res := mergr(arr)
	fmt.Printf("返回结果:%v\n", res)
}
