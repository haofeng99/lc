package pointer

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	l, r := 0, n-1
	for l <= r {
		if nums[l]+nums[r] < target {
			ans = ans + r - l
			l++
		} else {
			r--
		}
	}
	return ans
}

// https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
