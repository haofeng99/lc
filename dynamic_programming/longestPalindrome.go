package dynamicprogramming

// https://leetcode.cn/problems/longest-palindromic-substring/
// 5. 最长回文子串

// 动态规划法
// dp[l][r] 表示字符串s的子串s[l:r+1]是否是回文串
// 如果dp[l][r]为true, 要判断dp[l-1][r+1]是否为true, 只需要判断字符串在l-1和r+1位置的字符是否相等即可
// 矩阵初始化: 当l==r时, dp[l][r]为true, 因为单个字符一定是回文串;
// dp[l][r] = dp[l+1][r-1] && s[l] == s[r]
func longestPalindrome(s string) string {
	// 矩阵初始化
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	// 动态规划填表
	maxLen := 1
	start := 0
	// 为什么是先遍历j, 再遍历i, 因为dp[i][j]依赖于dp[i+1][j-1], 只有当j先遍历时, dp[i+1][j-1]才会被计算出来
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			// j-i<=2的判断条件为了防止dp[i+1][j-1]得到的结果为i'>j'的情况出现,
			// 此时dp[i'][j']的数据没有初始化,一定是false,所以计算结果就会被识别为false
			// 例如aa就会被认为不是回文串(dp[1][0]没有初始化,默认为false), 但是实际上aa是回文串
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					start = i
				}
			}
		}
	}
	return s[start : start+maxLen]
}

// 动态规划_内存优化
// 思路: // todo 理清优化思路
func longestPalindrome_ii(s string) string {
	// 矩阵初始化
	dp := make([]bool, len(s))

	// 动态规划填表
	maxLen := 1
	start := 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1]) {
				dp[i] = true
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					start = i
				}
			} else {
				dp[i] = false
			}
		}
	}
	return s[start : start+maxLen]
}

// 中心扩展法
// 核心思路:
// 回文串具有中心对称性, 即正向读和逆向读的结果都是一样的, 其中心可能是一个字符（如 "aba" 的中心是 'b'），也可能是两个字符（如 "abba" 的中心是 "bb"）
// 遍历字符串中的每个字符,将其作为回文串的中心,或者将其与下一个字符一起作为中心,向两边扩展,以找出以该中心为基准的最长回文子串。
// (一共就两种情况, 以一个字符为中心, 以两个字符为中心, 以三个字符为中心本质上还是以1个字符为中心)
func longestPalindrome_center_spread(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // 回文字符串中心是两一个字符
		len2 := expandAroundCenter(s, i, i+1) // 回文字符串中心是两个字符
		maxLen := max(len1, len2)
		// 针对奇数长度和偶数长度回文串的start和end的通用计算公式
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
