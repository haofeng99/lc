package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 123 买卖股票的最佳时机 III (买入卖出限制2笔)

// 思路: 由于限制了买入卖出次数, 因此需要在状态里加入交易次数的维度
// dp[i][j][k] 表示第i天, 交易了j次, 手里是否持有股票(k=0表示不持有, k=1表示持有)的最大利润
// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]) // 昨天没有持股而今天买入一只，故昨天买入的次数是j-1
func maxProfit_III(prices []int) int {
	// 矩阵初始化
	// 三维矩阵的解法不考虑
	return 0
}

// 思路: dp[i][j] 表示在第i天状态为j时的最大收益
// 其中j有以下5个状态
// 0: 未交易
// 1: 第一次交易的买入
// 2: 第一次交易的卖出
// 3: 第二次交易的买入
// 4: 第二次交易的卖出

func maxProfit_III_simple(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)

	// 矩阵初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0] // 需要考虑这种场景

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[n-1][4]
}

// 思考: 题目要求是最多可以进行2次交易, 为什么最后直接取的sell2的状态
// 观察上面的递推公式, 状态转移过程中, dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i]) // 这个过程, 如果dp[i-1][2]很小，可能就不进行买入了, 最终可能就是1次交易的结果

// 内存优化
func maxProfit_III_od(prices []int) int {
	no_op, buy_1, sell_1, buy2, sell_2 := 0, -prices[0], 0, -prices[0], 0
	for i := 0; i < len(prices); i++ {
		temp_no_op := no_op
		temp_buy_1 := max(buy_1, no_op-prices[i])
		temp_sell_1 := max(sell_1, buy_1+prices[i])
		temp_buy_2 := max(buy2, sell_1-prices[i])
		temp_sell_2 := max(sell_2, buy2+prices[i])

		no_op = temp_no_op
		buy_1 = temp_buy_1
		sell_1 = temp_sell_1
		buy2 = temp_buy_2
		sell_2 = temp_sell_2
	}
	return sell_2
}
