package sort

import "fmt"

// 315 计算右侧小于当前元素的个数
var ans []int
var extNums []extNum

type extNum struct {
	Value int
	Index int
}

func countSmaller(nums []int) []int {
	ans = make([]int, len(nums))
	extNums = []extNum{}
	for idx, value := range nums {
		extNums = append(extNums, extNum{value, idx})
	}

	splitCS(extNums)

	return ans

}

func splitCS(nums []extNum) []extNum {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := splitCS(nums[0:mid])
	right := splitCS(nums[mid:])

	return mergeCS(left, right)
}

func mergeCS(left, right []extNum) []extNum {
	m, n := len(left), len(right)
	merge := make([]extNum, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if left[i].Value <= right[j].Value {
			merge[k] = left[i]
			ans[left[i].Index] += j
			k++
			i++
		} else {
			merge[k] = right[j]
			k++
			j++
		}
	}

	if i == m {
		for j < n {
			merge[k] = right[j]
			k++
			j++
		}
	}
	if j == n {
		for i < m {
			merge[k] = left[i]
			ans[left[i].Index] += j
			k++
			i++
		}
	}

	return merge
}

func Show_countSmaller() {
	nums := []int{-1}
	fmt.Println(countSmaller(nums))
}

// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/description/
