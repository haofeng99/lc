package pointer

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	// 特判
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	idx := 0
	hasProcess := false
	intervals = append(intervals, newInterval)
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= newInterval[0] && !hasProcess {
			// i及其后面的元素统一向后移位, 同时i的位置留给newInterval
			for j := len(intervals) - 1; j > i; j-- {
				intervals[j] = intervals[j-1]
			}
			intervals[i] = newInterval
			hasProcess = true
		}
		if i >= 1 {
			if intervals[i][0] <= intervals[idx][1] {
				// 可以合并
				intervals[idx][1] = max(intervals[i][1], intervals[idx][1])
			} else {
				idx++
				intervals[idx] = intervals[i]
			}
		}
	}
	return intervals[:idx+1]
}

// https://leetcode.cn/problems/insert-interval/

func Show_insert() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	res := insert(intervals, newInterval)
	fmt.Println(res)
}
