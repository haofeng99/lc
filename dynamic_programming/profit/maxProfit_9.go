package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/
// 3573. 买卖股票的最佳时机 V
func maximumProfit(prices []int, k int) int64 {
	// dp[i][j][0] 表示第i天, 交易了j次, 手里不持有股票的最大利润
	// dp[i][j][1] 表示第i天, 交易了j次, 手里持有股票的最大利润
	// dp[i][j][2] 表示第i天, 交易了j次, 手里不持有股票的最大利润
	n := len(prices)
	dp := make([][][]int64, n)
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
		}
	}

	// 初始化第 0 天的状态
	for j := 1; j <= k; j++ {
		dp[0][j][1] = int64(-prices[0])
		dp[0][j][2] = int64(prices[0])
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j][1]+int64(prices[i]), dp[i-1][j][2]-int64(prices[i])))
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-int64(prices[i]))
			dp[i][j][2] = max(dp[i-1][j][2], dp[i-1][j-1][0]+int64(prices[i]))
		}
	}

	return dp[n-1][k][0]
}

func maximumProfit_(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][3]int64, k+1)
	// 初始化第 0 天的状态
	for j := 1; j <= k; j++ {
		dp[j][1] = int64(-prices[0])
		dp[j][2] = int64(prices[0])
	}

	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = max(dp[j][0], max(dp[j][1]+int64(prices[i]), dp[j][2]-int64(prices[i])))
			dp[j][1] = max(dp[j][1], dp[j-1][0]-int64(prices[i]))
			dp[j][2] = max(dp[j][2], dp[j-1][0]+int64(prices[i]))
		}
	}

	return dp[k][0]
}
