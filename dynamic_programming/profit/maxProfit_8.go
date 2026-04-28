package dynamicprogramming

import "math"

// 最少交易k次
// 思路: 将问题转化为刚好交易k次+无限次交易的子问题
// ans = max(刚好k次, 刚好k+1次, ...)
// 由于交易次数超过n/2后利润不再增加，上界取min(k+lim, n/2)

func maxProfit_VIII(prices []int, k int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return maxProfit_II(prices)
	}

	// 最多交易次数不会超过n/2，额外多试几次（利润递减后即停止有意义的上界）
	maxK := min(n/2, k+5)
	ans := math.MinInt32
	for t := k; t <= maxK; t++ {
		ans = max(ans, maxProfit_VII(prices, t))
	}
	return normalize(ans)
}

// 最少交易k次 直接DP
// dp[i][j][0] 表示前i天，至少交易j次，不持有股票
// dp[i][j][1] 表示前i天，至少交易j次，持有股票
// 关键: 买入时可以从dp[i-1][j-1][0]（刚好j-1次→至少j次）或dp[i-1][j][0]（已至少j次→继续交易）转移
func maxProfit_VIII_(prices []int, k int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return maxProfit_II(prices)
	}

	// j=0 等价于不限次数（maxProfit_II）
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}

	// 初始化
	dp[0][0][0] = 0
	dp[0][0][1] = -prices[0]
	for j := 1; j <= k; j++ {
		dp[0][j][0] = math.MinInt32
		dp[0][j][1] = math.MinInt32
	}
	dp[0][1][1] = -prices[0]

	for i := 1; i < n; i++ {
		// j=0: 不限次数
		dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][0][1]+prices[i])
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-prices[i])

		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 买入：从j-1（升到j）或从j（保持≥j）转移
			pre := dp[i-1][j-1][0]
			if dp[i-1][j][0] > pre {
				pre = dp[i-1][j][0]
			}
			dp[i][j][1] = dp[i-1][j][1]
			if pre != math.MinInt32 {
				dp[i][j][1] = max(dp[i][j][1], pre-prices[i])
			}
		}
	}

	return normalize(dp[n-1][k][0])
}

// 最少交易k次 状态压缩
func maxProfit_VIII_od(prices []int, k int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return maxProfit_II(prices)
	}

	dp := make([][2]int, k+1)
	for j := 1; j <= k; j++ {
		dp[j][0] = math.MinInt32
		dp[j][1] = math.MinInt32
	}
	dp[0][1] = -prices[0]
	dp[1][1] = -prices[0]

	for i := 1; i < n; i++ {
		// j=0 先算
		pre0 := dp[0][0]
		dp[0][0] = max(dp[0][0], dp[0][1]+prices[i])
		dp[0][1] = max(dp[0][1], pre0-prices[i])

		for j := k; j >= 1; j-- {
			// pre在dp[j][0]更新前取昨天值：max(昨天j-1次不持有, 昨天已≥j次不持有)
			pre := dp[j-1][0]
			if j == 1 {
				pre = pre0 // dp[0][0]已更新为今天的值，用昨天的pre0代替
			}
			if dp[j][0] > pre {
				pre = dp[j][0]
			}
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			if pre != math.MinInt32 {
				dp[j][1] = max(dp[j][1], pre-prices[i])
			}
		}
	}

	return normalize(dp[k][0])
}
