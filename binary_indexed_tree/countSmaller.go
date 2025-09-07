package binaryindexedtree

import (
	"fmt"
	"sort"
)

var a, c []int
var idxMap = make(map[int]int)

func countSmaller(nums []int) []int {
	ans := []int{}
	// 1.离散化
	discretization(nums)

	// 2.建树
	c = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		idx := idxMap[nums[i]]
		ans = append(ans, query(idx-1))
		update(idx)
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func lowbit(x int) int {
	return x & (-x)
}

// 单点更新
func update(pos int) {
	for pos < len(c) {
		c[pos]++
		pos += lowbit(pos)
	}
}

// 区间查询[从最开始查]
func query(pos int) int {
	ans := 0
	for pos > 0 {
		ans += c[pos]
		pos -= lowbit(pos)
	}
	return ans
}

// 数组离散化
func discretization(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	a = make([]int, 0, len(nums))
	for num := range set {
		a = append(a, num)
	}
	sort.Ints(a)

	for i, v := range a {
		idxMap[v] = i + 1
	}
}

func Show_countSmaller() {
	nums := []int{5, 2, 6, 1}
	ans := countSmaller(nums)
	fmt.Println(ans)
}

// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/description/
// 315. 计算右侧小于当前元素的个数
// 归并排序 + 树状数组
