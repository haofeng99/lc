package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
// 309. 买卖股票的最佳时机含冷冻期

// buy[i]  ：第 i 次买入后的最大利润
// sell[i] ：第 i 次卖出后的最大利润
// 只有发生了第i次买入, 才会有第i次卖出
// buy[i] = max(buy[i-1], sell[i-2]-prices[i]) // 什么都不做, 买入
// sell[i] = max(sell[i-1], buy[i-1] + prices[i]) // 什么都不做, 卖出
func maxProfit_V(prices []int) int {
	// if len(prices) == 1 {
	// 	return 0
	// }

	// // 这里直接是操作次数, 所以len(prices)
	// buy := make([]int, len(prices))
	// sell := make([]int, len(prices))

	// buy[0] = -prices[0]
	// sell[0] = 0

	// // 因为1-2为负数, 所以这里需要特殊处理
	// buy[1] = max(buy[0], 0-prices[1]) // sell[-1]等于0
	// sell[1] = max(sell[0], buy[0]+prices[1])

	// for i := 2; i < len(prices); i++ {
	// 	buy[i] = max(buy[i-1], sell[i-2]-prices[i])
	// 	sell[i] = max(sell[i-1], buy[i-1]+prices[i])

	// }

	// return sell[len(prices)-1]
	return 0
}

func maxProfit_V_(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}
func maxProfit_V__(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2-prices[i])
		newf1 := f0 + prices[i]
		newf2 := max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}
