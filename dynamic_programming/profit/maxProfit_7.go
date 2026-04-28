package dynamicprogramming

import "math"

// 刚好交易K次
// 与"最多k笔"(maxProfit_IV)的区别：初始化时将不可达状态设为-INF
// 只有dp[0][0][0]=0（第0天0次交易不持有）和dp[0][1][1]=-prices[0]（第0天买入1次）可达
// 最终答案在dp[n-1][k][0]：恰好完成k次交易且不持有股票
func maxProfit_VII(prices []int, k int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j][0] = math.MinInt32
			dp[i][j][1] = math.MinInt32
		}
	}

	dp[0][0][0] = 0
	dp[0][1][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0][0] = dp[i-1][0][0]
		for j := 1; j <= k; j++ {
			// 今天不持有：昨天不持有 或 昨天持有今天卖出（卖出不改变交易次数）
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 今天持有：昨天持有 或 昨天不持有(j-1次)今天买入（买入增加交易次数）
			dp[i][j][1] = dp[i-1][j][1]
			if dp[i-1][j-1][0] != math.MinInt32 {
				dp[i][j][1] = max(dp[i][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}

	return normalize(dp[n-1][k][0])
}

// 状态压缩
func maxProfit_VII_(prices []int, k int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	dp := make([][2]int, k+1)
	for j := 0; j <= k; j++ {
		dp[j][0] = math.MinInt32
		dp[j][1] = math.MinInt32
	}
	dp[0][0] = 0
	dp[1][1] = -prices[0]

	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			if dp[j-1][0] != math.MinInt32 {
				dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
			}
		}
	}

	return normalize(dp[k][0])
}

func normalize(v int) int {
	if v == math.MinInt32 {
		return 0
	}
	return v
}
