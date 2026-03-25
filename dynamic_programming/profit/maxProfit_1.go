package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 121. 买卖股票的最佳时机 (最多只买卖一次)

// dp[i]表示第i天卖出获得的最大利润
// dp[i] = max(price[i]-i_pre_min, dp[i-1])
// 如果第i天卖出股票，则最大利润为(该天的股价-前面天数中最小的股价)，然后与已知的最大利润比较，如果大于则更新当前最大利润的值
func maxProfit(prices []int) int {
	minPrice := prices[0]
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		dp[i] = max(prices[i]-minPrice, dp[i-1])
	}
	return dp[len(dp)-1]
}

// 思路: 找最小值和最大值
func maxProfit_(prices []int) int {
	minPrice := prices[0]
	curMaxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		curMaxProfit = max(curMaxProfit, prices[i]-minPrice)
	}
	return curMaxProfit
}

// 贪心算法
