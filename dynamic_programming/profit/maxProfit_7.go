package dynamicprogramming

// 刚好交易K次
// 思路: DP, f[i][j][0] 表示前i天，恰好进行了j-1次交易，不持有股票的最大利润
// j从1到k+1，其中j=1表示0次交易，j=k+1表示k次交易
func maxProfit_VII(prices []int, k int) int {
	const INF = 1e9
	n := len(prices)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, k+2)
		for j := range f[i] {
			f[i][j] = make([]int, 2)
			for h := range f[i][j] {
				f[i][j][h] = -INF
			}
		}
	}
	f[0][1][0] = 0
	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
		}
	}
	return f[n][k+1][0]
}

// 状态压缩

func maxProfit_VII_(prices []int, k int) int {
	const INF = 1e9
	n := len(prices)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	for j := 0; j <= k; j++ {
		dp[0][j][0] = -INF
		dp[0][j][1] = -INF
	}

	dp[0][1][0] = 0
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}
