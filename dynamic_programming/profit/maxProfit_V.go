package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
// 309. 买卖股票的最佳时机含冷冻期

// buy[i]  ：第 i 次买入后的最大利润
// sell[i] ：第 i 次卖出后的最大利润
// 只有发生了第i次买入, 才会有第i次卖出
// buy[i] = max(buy[i-1], sell[i-2]-prices[i]) // 什么都不做, 买入
// sell[i] = max(sell[i-1], buy[i-1] + prices[i]) // 什么都不做, 卖出
func maxProfit_V(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	buy := make([]int, len(prices))
	sell := make([]int, len(prices))

	buy[0] = -prices[0]
	sell[0] = 0

	buy[1] = max(buy[0], 0-prices[1]) // sell[-1]等于0
	sell[1] = max(sell[0], buy[0]+prices[1])

	for i := 2; i < len(prices); i++ {
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])

	}

	return sell[len(prices)-1]
}
