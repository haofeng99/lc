package divideconquer

import "strings"

// 395. 至少有 K 个重复字符的最长子串
// 思路: 统计每一个字符出现的次数, 如果一个字符出现的次数小于k, 则所有包含k的子串都不满足条件
// 因此可以针对该字符进行源字符串分割, 然后分治求解

// 时间复杂度O(n*26*26) 空间复杂度 O(26*26)
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	charCntMap := make(map[rune]int)
	for _, char := range s {
		charCntMap[char] += 1
	}

	for char, cnt := range charCntMap {
		if cnt < k {
			var ans = 0
			sibStrArr := strings.Split(s, string(char)) // 时间复杂度的n主要体现在这里
			for _, subStr := range sibStrArr {
				ans = max(ans, longestSubstring(subStr, k))
			}
			return ans
		}
	}
	return len(s)
}

// 另外一个思路: 滑动窗口
