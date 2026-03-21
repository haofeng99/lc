package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
// 714. 买卖股票的最佳时机含手续费
// 不限次数买卖扣费场景
// 关键点: 将扣费过程统一, 要不买入扣费, 要不卖出扣费
func maxProfit_VI(prices []int, fee int) int {
	hold, sell := -prices[0]-fee, 0
	for i := 1; i < len(prices); i++ {
		hold_pre, sell_pre := hold, sell
		hold = max(hold_pre, sell_pre-prices[i]-fee)
		sell = max(hold_pre+prices[i], sell_pre)
	}
	return max(hold, sell)
}

// 卖出扣费
// 需要注意的点: 初始化时, 卖出不扣费
// 假设 prices = [1]，fee = 1：
// 原来的代码：hold = -1, sell = -1 → 返回 max(-1, -1) = -1（错误，因为无法交易，利润应为 0）
func maxProfit_VI_1(prices []int, fee int) int {
	hold, sell := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		hold_pre, sell_pre := hold, sell
		hold = max(hold_pre, sell_pre-prices[i])
		sell = max(hold_pre+prices[i]-fee, sell_pre)
	}
	return max(hold, sell)
}
