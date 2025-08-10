package priorityqueue

import "sort"

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	return 0
}

// https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/
