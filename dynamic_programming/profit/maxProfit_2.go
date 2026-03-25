package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 122 买卖股票的最佳时机 II (买入卖出不限次数)

// dp动态规划 思路: 某一天两个状态, 持有和卖出
// dp[i][0] 表示第i天手里没有股票最高收益 dp[i][1] 表示第i天手里有股票最高收益
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - price[i])
// 初始化
func maxProfit_II(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	// 初始化
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

// 内存优化
func maxProfit_II_od(prices []int) int {
	hold, sell := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		hold_pre, sell_pre := hold, sell
		hold = max(hold_pre, sell_pre-prices[i])
		sell = max(hold_pre+prices[i], sell_pre)
	}
	return max(hold, sell)
}

// 非dp做法
// 思路: 由于不限制购买和卖出次数, 所以最简单的就是直接吃到所有涨幅, 避开所有跌幅
// 从头开始遍历, 步长为1, 只要是涨幅, 就算利润里
func maxProfit_II_greedy(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
