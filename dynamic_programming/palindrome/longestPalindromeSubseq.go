package dynamicprogramming

// https://leetcode.cn/problems/longest-palindromic-subsequence/
// 516. 最长回文子序列

// 动态规划
// dp[i][j]表示i到j范围内可以组成的最长回文串的长度
func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1 // 必须要做初始化
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ { // 为什么j要从i+1开始, 防止下面数组越界
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]
}

// 内存优化
// 内层循环过程中, i的值是锁定的, 因此可以节省一个维度
// 脑中构思二维数组的计算流程, 每个位置的计算需要哪几个方向上的数据, 如果没有某个方向上的数据,就需要定义额外的变量来存储这个数据
func longestPalindromeSubseq_ii(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = 1 // 必须要做初始化
		deep_pre := 0
		for j := i + 1; j < len(s); j++ {
			deep_pre_snapshot := dp[j]
			if s[i] == s[j] {
				dp[j] = deep_pre + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			deep_pre = deep_pre_snapshot
		}
	}

	return dp[len(s)-1]
}
