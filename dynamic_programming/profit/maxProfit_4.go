package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 188. 买卖股票的最佳时机 IV (买入卖出限制最多k笔)

// 思路:
// buy[i]  ：第 i 次买入后的最大利润
// sell[i] ：第 i 次卖出后的最大利润
// 只有发生了第i次买入, 才会有第i次卖出
// buy[i] = max(buy[i], sell[i-1]-prices[i]) // 什么都不做, 买入
// sell[i] = max(sell[i], buy[i] + prices[i]) // 什么都不做, 卖出
func maxProfit_IV(k int, prices []int) int {

	// // 关键点: 为什么需要设置为k+1的长度
	// // buy和sell定义为第i次操作的最大利润, 一共需要操作1->k次
	// // 同时因为buy和sell需要依赖前一步的状态, buy[i]以来sell[i-1], 所以需要多一个长度
	// buy := make([]int, k+1)
	// sell := make([]int, k+1)

	// // 关键点: 为什么buy[i]需要设置为-price[0]或者负无穷
	// // 首先, buy[i]绝对不能默认是0, 否则就会是0元买股票的情况
	// // 其次, 直接看for循环过程, 内层循环完全是在循环买卖次数, 所以最初时, buy[i]一定是从-price[0]开始的
	// for i := 1; i <= k; i++ {
	// 	buy[i] = math.MinInt
	// }

	// for i := 0; i < len(prices); i++ {
	// 	for j := 1; j <= k; j++ {
	// 		buy[j] = max(buy[j], sell[j-1]-prices[i])
	// 		sell[j] = max(sell[j], buy[j]+prices[i])
	// 	}
	// }

	// return sell[k]

	// 非主流解法 遗忘
	return 0
}

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

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]
}
