package hash

// https://leetcode.cn/problems/first-missing-positive/
// 41. 缺失的第一个正数

func firstMissingPositive(nums []int) int {
	n := len(nums)
	ans := 1
	for _, num := range nums {
		if num >= n {
			continue
		}

	}
	return ans
}
