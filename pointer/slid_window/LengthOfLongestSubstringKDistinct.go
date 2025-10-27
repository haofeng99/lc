package slidwindow

// https://www.lintcode.com/problem/386/solution?showListFe=true&page=8&ordering=id&pageSize=50
// leetcode 340
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}
	left, right := 0, 1
	charCntMap := map[rune]int{
		rune(s[0]): 1,
	}

	ans := 1

	for left <= right && right < len(s) {
		charCntMap[rune(s[right])]++

		for len(charCntMap) > k && left <= right {
			leftChar := rune(s[left])
			charCntMap[leftChar]--
			if charCntMap[leftChar] == 0 {
				delete(charCntMap, leftChar)
			}
			left++
		}

		ans = max(ans, right-left+1)

		right++
	}
	return ans
}
