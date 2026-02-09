package dynamicprogramming

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 122 买卖股票的最佳时机 II
func maxProfit_II(prices []int) int {
	min_price := prices[0]
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1]+prices[i], min_price)
	}
	return dp[len(dp)-1]
}
