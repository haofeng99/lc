package dynamicprogramming

// https://leetcode.cn/problems/palindromic-substrings/
// 647. 回文子串
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	cnt := 0
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				cnt++
			}
		}
	}
	return cnt
}

// 动态规划内存优化
func countSubstrings_ii(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]bool, len(s))

	cnt := 0
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1]) {
				dp[i] = true
				cnt++
			} else {
				dp[i] = false
			}
		}
	}
	return cnt
}

// 中心扩散法1_奇数偶数长度中心
func countSubstrings_center_spread(s string) int {
	if len(s) == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt1 := count_expandAroundCenter(s, i, i)
		cnt2 := count_expandAroundCenter(s, i, i+1)
		cnt += cnt1 + cnt2
	}
	return cnt
}
func count_expandAroundCenter(s string, i, j int) int {
	ans := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		ans++
	}
	return ans
}

// 中心扩展法2_中心个数2*len(s)-1
// 为什么是这个个数:https://leetcode.cn/problems/palindromic-substrings/solutions/154773/liang-dao-hui-wen-zi-chuan-de-jie-fa-xiang-jie-zho/
func countSubstrings_center_spread_ii(s string) int {
	if len(s) == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < 2*len(s)-1; i++ {
		// i/2表示中心位置, i%2表示中心是一个字符还是两个字符之间的空隙
		// 例如: 当i=0时, l=r=0, 中心是第一个字符; 当i=1时, l=0, r=1, 中心是第一个字符和第二个字符;
		// 当i=2时, l=r=1, 中心是第二个字符; 以此类推
		l, r := i/2, i/2+i%2
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			cnt++
		}
	}
	return cnt
}
