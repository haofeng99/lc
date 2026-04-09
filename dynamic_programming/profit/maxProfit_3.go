package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 123 买卖股票的最佳时机 III (买入卖出限制2笔)

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
	no_op := 0
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0

	for i := 1; i < len(prices); i++ {
		pre_no_op, pre_buy1, pre_sell1, pre_buy2, pre_sell2 := no_op, buy1, sell1, buy2, sell2

		no_op = pre_no_op
		buy1 = max(pre_buy1, pre_no_op-prices[i])
		sell1 = max(pre_sell1, pre_buy1+prices[i])
		buy2 = max(pre_buy2, pre_sell1-prices[i])
		sell2 = max(pre_sell2, pre_buy2+prices[i])
	}

	return sell2
}
