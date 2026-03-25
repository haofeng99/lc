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
	// dp[i][j][0] 表示第i天, 交易了j次, 手里不持有股票的最大利润
	// dp[i][j][1] 表示第i天, 交易了j次, 手里持有股票的最大利润
	n := len(prices)
	f := make([][][2]int, n)
	for i := range f {
		f[i] = make([][2]int, k+1)
		for j := range f[i] {
			f[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2} // 防止溢出
		}
	}
	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}
	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
		}
	}
	return f[n][k+1][0]
}
func maxProfit__(k int, prices []int) int {
	f := make([][2]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2 // 防止溢出
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
		}
	}
	return f[k+1][0]
}
