package dynamicprogramming

// https://leetcode.cn/problems/longest-increasing-subsequence
// 300 最长递增子序列
// 扩展题
// 1. 输出子序列的路径
// 2. 字典序最小的路径

// 动态规划
// dp[i]表示以第i个数字结尾的最长上升子序列的长度
// dp[i] = dp[j]+1 (nums[i] > nums[j] && i > j)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = 1
	maxans := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
			maxans = max(maxans, dp[i])
		}
	}
	return maxans
}
