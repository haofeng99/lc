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
	}

	// 动态规划填表
	maxLen := 1
	start := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ { // dp矩阵构造时没有初始化值,j从i开始计算
			// j-i <= 1 是为了防止一些数值没有初始化的情况
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
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
// 思路:
// 在内层循环中, 此时dp[i][j]的j值已经是固定的了, 因此没必要再去维护, 因此只需要计算当前j列的值就行, 并且把当前j列计算出的结果作为下一个j列的计算基础;
// 但是需要注意的是, 如果当前j列某个位置不满足条件, 需要手动把对应位置设置为false, 否则可能因为历史数据记录影响后续的计算
func longestPalindrome_ii(s string) string {
	// 矩阵初始化
	dp := make([]bool, len(s))

	maxLen := 1
	start := 0
	for i := len(s) - 1; i >= 0; i-- {
		deep_pre := false
		for j := i; j < len(s); j++ {
			temp := dp[j]
			if s[i] == s[j] && (j-i <= 1 || deep_pre) {
				dp[j] = true
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					start = i
				}
			} else {
				dp[j] = false
			}
			deep_pre = temp
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

// 中心扩张法2: 以2*len(s)-1个中心进行扩张
func longestPalindrome_center_spread_ii(s string) string {
	if len(s) == 0 {
		return ""
	}
	max_len := 0
	left := 0
	for i := 0; i < 2*len(s)-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > max_len {
				max_len = r - l + 1
				left = l
			}
			l--
			r++
		}
	}
	return s[left : left+max_len]
}
