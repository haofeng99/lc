package sort

import "sort"

func frequencySort(nums []int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		if cnt[a] == cnt[b] {
			return a > b
		}
		return cnt[a] < cnt[b]
	})
	return nums
}

// https://leetcode.cn/problems/sort-array-by-increasing-frequency/
