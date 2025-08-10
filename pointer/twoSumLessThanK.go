package pointer

import (
	"fmt"
	"sort"
)

// 排序 + 双指针
func twoSumLessThanK(vector []int, k int) int {
	sort.Ints(vector)
	ans, min := -1, 1000
	for i := 0; i < len(vector); i++ {
		if vector[i] >= k {
			return ans
		}
		for j := i + 1; j < len(vector); j++ {
			if vector[i]+vector[j] > k {
				break
			}
			if abs(vector[i]+vector[j]-k) < min {
				min = abs(vector[i] + vector[j] - k)
				ans = vector[i] + vector[j]
			}
		}
	}
	return ans
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

// TODO 排序 + 二分查找

func Show_twoSumLessThanK() {
	vector := []int{10, 20, 30}
	ans := twoSumLessThanK(vector, 15)
	fmt.Println(ans)
}
