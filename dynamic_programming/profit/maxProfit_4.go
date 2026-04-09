package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 188. 买卖股票的最佳时机 IV (买入卖出限制最多k笔)

func maxProfit_IV_(k int, prices []int) int {
	// 状态定义
	// dp[i][j][0] 表示第i天, 交易了j次, 手里不持有股票的最大利润
	// dp[i][j][1] 表示第i天, 交易了j次, 手里持有股票的最大利润

	// 状态转移
	// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
	// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])

	// 矩阵初始化
	// 第0天, 交易j次, 手里不持有股票的最大利润为0
	// 第0天, 交易j次, 手里持有股票的最大利润为-price[0] (需要考虑这种情况)

	n := len(prices)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	for j := 1; j <= k; j++ {
		dp[0][j][0] = 0
		dp[0][j][1] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

// 按天压缩
func maxProfit__(k int, prices []int) int {
	n := len(prices)
	dp := make([][2]int, k+1)
	for j := 1; j <= k; j++ {
		dp[j][0] = 0
		dp[j][1] = -prices[0]
	}

	// 为什么内层需要逆序便利, 想象二维矩阵的多次印刷
	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]
}
