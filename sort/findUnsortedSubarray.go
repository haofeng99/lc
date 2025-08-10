package sort

import "sort"

// 排序后前缀和后缀子数组都不会变
func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	copyNums := append([]int{}, nums...)
	sort.Ints(copyNums)
	left := 0
	for left < len(nums) {
		if nums[left] != copyNums[left] {
			break
		}
		left++
	}
	right := len(nums) - 1
	for right >= 0 {
		if nums[right] != copyNums[right] {
			break
		}
		right--
	}
	if left != right && left < right {
		return right - left + 1
	}
	return 0
}

// 将整个数组分成三段，左边是递增的，中间是乱序的，右边是递增的
// 从左向右看:如果进入了右段，就没有比最大值小的数，所以最后一个比最大值小的数就是中段的右边界
// 从右向左看:如果进入了左段，就没有比最小值大的数，所以最后一个比最小值大的数就是中段的左边界
func findUnsortedSubarray2(nums []int) int {
	left, right := 0, len(nums)
	min, max := nums[0], nums[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}

		j := len(nums) - 1 - i
		if nums[j] > min {
			left = j
		} else {
			min = nums[j]
		}
	}
	if left < right {
		return right - left + 1
	} else {
		return 0
	}
}

// https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/
