package divideconquer

import "math"

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 4. 寻找两个正序数组的中位数
// 思路: 找第k小的数据

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保nums1是较短的数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		// partition 表示左边有多少个元素
		partition1 := (low + high) / 2
		partition2 := (m+n+1)/2 - partition1

		// 处理边界情况，获取四个关键值
		maxLeft1 := math.MinInt32
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		minRight1 := math.MaxInt32
		if partition1 < m {
			minRight1 = nums1[partition1]
		}

		maxLeft2 := math.MinInt32
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		minRight2 := math.MaxInt32
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		// 检查划分是否正确
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 1 {
				return float64(max(maxLeft1, maxLeft2))
			}
			return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
		} else if maxLeft1 > minRight2 {
			high = partition1 - 1
		} else {
			low = partition1 + 1
		}
	}

	return 0.0
}

// 为什么使用下标法无法解决这道题(当partition1为1时表示左边有2个元素)
// 1.使用下标法时,在分割数组时,需要支持单边不存在元素,因此下标需要满足[-1, len(nums)-1], 很显然这是行不通的
// 2.因此建议partition不表示下标,而标识左边元素的个数
func findMedianSortedArraysII(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArraysII(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	// 特判
	if m == 0 {
		if n%2 == 1 {
			return float64(nums2[n/2])
		} else {
			return float64(nums2[n/2]+nums2[n/2-1]) / 2.0
		}
	}

	left, right := 0, m-1
	for left <= right {
		partition1 := left + (right-left)/2
		partition2 := (m+n+1)/2 - (partition1 + 1) - 1

		leftMax1 := math.MinInt32
		rightMin1 := math.MaxInt32
		leftMax2 := math.MinInt32
		rightMin2 := math.MaxInt32

		if partition1 >= 0 {
			leftMax1 = nums1[partition1]
		}
		if partition1 < m-1 {
			rightMin1 = nums1[partition1+1]
		}
		if partition2 >= 0 {
			leftMax2 = nums2[partition2]
		}
		if partition2 < n-1 {
			rightMin2 = nums2[partition2+1]
		}

		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			// 找到结果了
			if (m+n)%2 == 1 {
				return float64(max(leftMax1, leftMax2))
			}
			return float64(max(leftMax1, leftMax2)+min(rightMin1, rightMin2)) / 2.0
		} else if leftMax1 > rightMin2 {
			right = partition1 - 1
		} else {
			left = partition1 + 1
		}
	}

	return 0.0
}
