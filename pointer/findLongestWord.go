package pointer

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		}
		return dictionary[i] < dictionary[j]
	})

	n := len(s)
	for _, word := range dictionary {
		if len(word) > n {
			continue
		}
		i := 0
		for j := 0; j < n; j++ {
			if s[j] == word[i] {
				i++
				if i == len(word) {
					return word
				}
			}
		}
	}
	return ""
}

// https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/

func Show_findLongestWord() {
	dictionary := []string{"ba", "ab", "a", "b"}
	s := "bab"
	ans := findLongestWord(s, dictionary)
	fmt.Println(ans)
}

// test
func Show_findLongestWord2() {
	s := []int{2, 4, 5, 6}
	fmt.Printf("s: %p\n", s)
	s = s[1:]
	fmt.Printf("s: %p\n", s)

	charIdxSee := make(map[rune]int)
	charIdxSee['a'] += 1
	charIdxSee['b']++
	fmt.Println(charIdxSee)
	ok, num := charIdxSee['c']
	fmt.Println(ok, num)
}
