package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 122 买卖股票的最佳时机 II

// dp动态规划
// dp[i][0] 表示第i天手里没有股票最高收益 dp[i][1] 表示第i天手里有股票最高收益
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - price[i])
// 初始化
func maxProfit_II(prices []int) int {
	return 0
}

// 内存优化
func maxProfit_II_od(prices []int) int {
	min_price := prices[0]
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1]+prices[i], min_price)
	}
	return dp[len(dp)-1]
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
