package pointer

import "sort"

// 259：3Sum Smaller 三数据之和小于target
// 核心思路:
// 1.先排序数组
// 2.固定第一个数i
// 3.对区间 (i+1, n-1) 使用 双指针:
//
//	left = i + 1
//	right = n - 1
//	如果 nums[i] + nums[left] + nums[right] < target
//	  那么:left到right之间所有的k都满足一次性加right-left -> left++
//	否则 right--
func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < target {
				ans += right - left
				left++
			} else {
				right--
			}
		}
	}
	return ans
}
