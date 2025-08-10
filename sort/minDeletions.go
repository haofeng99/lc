package sort

import (
	"fmt"
	"sort"
)

func minDeletions(s string) int {
	chatCnt := make([]int, 26)
	for _, c := range s {
		chatCnt[c-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(chatCnt)))

	ans := 0
	for i := 1; i < len(chatCnt); i++ {
		for chatCnt[i] >= chatCnt[i-1] && chatCnt[i] > 0 {
			ans++
			chatCnt[i]--
		}
	}
	return ans
}

// https://leetcode.cn/problems/minimum-deletions-to-make-character-frequencies-unique/description/
func Show_minDeletions() {
	fmt.Println(minDeletions("bbcebab"))
}
