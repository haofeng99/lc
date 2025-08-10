package sort

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {
	pieceMap := make(map[int][]int)
	for _, v := range pieces {
		pieceMap[v[0]] = v
	}
	for i := 0; i < len(arr); {
		v, ok := pieceMap[arr[i]]
		if !ok {
			return false
		}
		for j := 0; j < len(v); j++ {
			if i+j > len(arr) || arr[i+j] != v[j] {
				return false
			}
		}
		delete(pieceMap, arr[i])
		i = i + len(v)
	}
	return true
}

// https://leetcode.cn/problems/check-array-formation-through-concatenation/

func Show_canFormArray() {
	arr := []int{91, 4, 64, 78}
	pieces := [][]int{
		{78},
		{4, 64},
		{91},
	}
	can := canFormArray(arr, pieces)
	fmt.Println(can)
}
