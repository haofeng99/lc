package binaryindexedtree

import (
	"fmt"
	"sort"
)

var aI, cI []int
var idxMapI = make(map[int]int)

func reversePairs(record []int) int {
	ans := 0
	// 1.离散化
	discretizationI(record)

	// 2.建树
	cI = make([]int, len(record))
	for i := len(record) - 1; i >= 0; i-- {
		idx := idxMapI[record[i]]
		ans += queryI(idx - 1)
		updateI(idx)
	}
	return ans
}

func lowbitI(x int) int {
	return x & (-x)
}

// 单点更新
func updateI(pos int) {
	for pos < len(cI) {
		cI[pos]++
		pos += lowbitI(pos)
	}
}

// 区间查询[从最开始查]
func queryI(pos int) int {
	ans := 0
	for pos > 0 {
		ans += cI[pos]
		pos -= lowbitI(pos)
	}
	return ans
}

// 数组离散化
func discretizationI(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	aI = make([]int, 0, len(nums))
	for num := range set {
		aI = append(aI, num)
	}
	sort.Ints(aI)

	for i, v := range aI {
		idxMapI[v] = i + 1
	}
}

func Show_reversePairs() {
	record := []int{9, 7, 5, 4, 6}
	ans := reversePairs(record)
	fmt.Println(ans)
}

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
