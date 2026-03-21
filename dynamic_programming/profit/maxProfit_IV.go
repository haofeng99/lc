package dynamicprogramming

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 188. 买卖股票的最佳时机 IV (买入卖出限制最多k笔)

// 思路:
// buy[i]  ：第 i 次买入后的最大利润
// sell[i] ：第 i 次卖出后的最大利润
// 只有发生了第i次买入, 才会有第i次卖出
// buy[i] = max(buy[i], sell[i-1]-prices[i]) // 什么都不做, 买入
// sell[i] = max(sell[i], buy[i] + prices[i]) // 什么都不做, 卖出
func maxProfit_IV(k int, prices []int) int {

	// 关键点: 为什么需要设置为k+1的长度
	// buy和sell定义为第i次操作的最大利润, 一共需要操作1->k次
	// 同时因为buy和sell需要依赖前一步的状态, buy[i]以来sell[i-1], 所以需要多一个长度
	buy := make([]int, k+1)
	sell := make([]int, k+1)

	// 关键点: 为什么buy[i]需要设置为-price[0]或者负无穷
	// 首先, buy[i]绝对不能默认是0, 否则就会是0元买股票的情况
	// 其次, 直接看for循环过程, 内层循环完全是在循环买卖次数, 所以最初时, buy[i]一定是从-price[0]开始的
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt
	}

	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	return sell[k]
}
