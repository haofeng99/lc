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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
