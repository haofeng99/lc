package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
// 309. 买卖股票的最佳时机含冷冻期

func maxProfit_V(prices []int) int {
	// 因为是在卖出股票后才会有冷静期, 所以需要针对不持有股票增加状态
	// dp[i][0] 第i天持有股票的最大收益
	// dp[i][1] 第i天不持有股票的最大收益 当前在冷冻期
	// dp[i][2] 第i天不持有股票的最大收益 当前不在冷冻期

	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	// 状态转移
	// dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
	// dp[i][1] = dp[i-1][0] + prices[i]
	// dp[i][2] = max(dp[i-1][1], dp[i-1][2])

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func maxProfit_V__(prices []int) int {
	dp0 := -prices[0]
	dp1 := 0
	dp2 := 0

	for i := 1; i < len(prices); i++ {
		pre_dp0, pre_dp1, pre_dp2 := dp0, dp1, dp2
		dp0 = max(pre_dp0, pre_dp2-prices[i])
		dp1 = pre_dp0 + prices[i]
		dp2 = max(pre_dp1, pre_dp2)
	}

	return max(dp1, dp2)
}
