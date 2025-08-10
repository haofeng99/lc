package sort

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	ans := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] == 0 {
			continue
		}
		if i >= 1 && citations[i] == citations[i-1] {
			continue
		}
		if len(citations)-i >= citations[i] {
			ans = max(ans, citations[i])
		} else {
			ans = max(ans, len(citations)-i)
		}

	}
	return ans
}

func hIndex2(citations []int) int {
	sort.Ints(citations)

	ans := 0
	for i := len(citations) - 1; i >= 0 && citations[i] > ans; i-- {
		ans++
	}
	return ans
}

// https://leetcode.cn/problems/h-index/
