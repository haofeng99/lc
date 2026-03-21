package enum

// 3871. 统计范围内的逗号 II
// https://leetcode.cn/problems/count-commas-in-range-ii/
// 枚举逗号
// 最右边的逗号 出现在 10^3 - n 之间  数量 n - 10^3 + 1
// 倒数第二个逗号 出现在 10^6 - n 之间 数量 n - 10^6 + 1
func countCommas(n int64) int64 {
	if n < 1000 {
		return 0
	}
	var ans int64 = 0
	for i := int64(1000); i <= n; i *= 1000 {
		ans += n - i + 1
	}
	return ans
}
