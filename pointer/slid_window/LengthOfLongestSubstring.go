package slidwindow

// https://www.lintcode.com/problem/384
// 最长无重复字符子串
func LengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	valIdxMap := make(map[uint8]int)
	ans := 0

	for l <= r && r < len(s) {
		char := s[r]
		if idx, ok := valIdxMap[char]; ok && idx >= l {
			l = idx + 1
		}
		valIdxMap[char] = r
		if r-l+1 > ans {
			ans = r - l + 1
		}
		r++
	}

	return ans
}
