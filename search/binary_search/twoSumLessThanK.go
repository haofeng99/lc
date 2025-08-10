package binarysearch

import (
	"fmt"
	"sort"
)

// TODO
func twoSumLessThanK(vector []int, k int) int {
	sort.Ints(vector)
	ans := 1
	for i := 0; i < len(vector); i++ {
		if vector[i] >= k {
			return ans
		}
		// target := l - vector[i]
	}
	return 0
}

// 在已经升序排序的切片nums的范围[i, j]内, 查找接近target的元素下标
func range_binary_search(nums []int, target int, i, j int) int {
	if nums[i] <= target {
		return i
	}
	if nums[j] >= target {
		return j
	}

	return 0
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

func Show_twoSumLessThanK() {
	vector := []int{10, 20, 30}
	ans := twoSumLessThanK(vector, 15)
	fmt.Println(ans)
}
