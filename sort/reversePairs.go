package sort

import "fmt"

var count int

func reversePairs(record []int) int {
	count = 0
	splitPairs(record)
	return count
}

func splitPairs(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	left := splitPairs(nums[0:mid])
	right := splitPairs(nums[mid:])

	return mergeAndCount(left, right)
}

func mergeAndCount(left, right []int) []int {
	m, n := len(left), len(right)
	res := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if left[i] <= right[j] {
			res[k] = left[i]
			// 此时left[i]是一定大于right[j]所有左边的数的
			// 这个是可以保证的
			count += j
			i = i + 1
		} else {
			res[k] = right[j]
			j = j + 1
		}
		k = k + 1
	}

	if i == m {
		for j < n {
			res[k] = right[j]
			j = j + 1
			k = k + 1
		}
	}
	if j == n {
		for i < m {
			res[k] = left[i]
			count += j
			i = i + 1
			k = k + 1
		}
	}
	return res
}

func Show_reversePairs() {
	nums := []int{4, 5, 6, 7}
	fmt.Println(reversePairs(nums))
}

// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/description/
// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
