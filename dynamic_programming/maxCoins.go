package dynamicprogramming

// 312. 戳气球

// 思路
// 1. 针对题目定义, 合理扩充数组 1, x,x,x, 1, 对最终结果不会造成影响
// 2. dp[i][j] 定义为戳破开区间i到j之内的所有气球得到的最大分数
// 3. 定义状态转移方程
// 3.1 dp[i][j] = dp[i][k] + val[i]*val[k]*val[j] + dp[k][j]
func maxCoins(nums []int) int {
	return 0
}
