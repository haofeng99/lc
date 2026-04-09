package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/
// 3573. 买卖股票的最佳时机 V
func maximumProfit(prices []int, k int) int64 {
	// dp[i][k][0] 第i天在交易k次后不持有股票的收益
	// dp[i][k][1] 第i天在交易k次后, 持有一只股票
	// dp[i][k][2] 第i天在交易k次后, 持有一只股票的空头

	n := len(prices)
	dp := make([][][3]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][3]int, k+1)
	}

	// dp[0][k][0] = 0
	// dp[0][k][1] = -prices[0]
	// dp[0][k][2] = prices[0]

	for i := 1; i <= k; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
		dp[0][i][2] = prices[0]
	}

	// 状态转移
	// dp[i][k][0] = max(dp[i-1][k][0], max(dp[i-1][k][1]+prices[i], dp[i-1][k][2]-prices[i]))
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
	// dp[i][k][2] = max(dp[i-1][k][2], dp[i-1][k-1][0]+prices[i])

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j][1]+prices[i], dp[i-1][j][2]-prices[i]))
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			dp[i][j][2] = max(dp[i-1][j][2], dp[i-1][j-1][0]+prices[i])
		}
	}

	return int64(dp[n-1][k][0])
}

func maximumProfit_(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][3]int, k+1)

	for i := 1; i <= k; i++ {
		dp[i][1] = -prices[0]
		dp[i][2] = prices[0]
	}

	// 为什么要内部逆序便利, 想像成一个二维矩阵的印刷, 每次刷新完要走新的一页继续刷新时, 想要使用历史的数据必须从后往前刷,
	// 如果从前往后刷, 那么j-1对应的其实是今天的j-1的数据, 而不是昨天的j-1的数据
	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			dp[j][0] = max(dp[j][0], max(dp[j][1]+prices[i], dp[j][2]-prices[i]))
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
			dp[j][2] = max(dp[j][2], dp[j-1][0]+prices[i])
		}
	}

	return int64(dp[k][0])
}
