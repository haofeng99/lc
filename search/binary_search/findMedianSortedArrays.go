package binarysearch

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 4. 寻找两个正序数组的中位数
// 思路: 找第k小的数据

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var findKth func([]int, []int, int) int
	// k是第几个最小数, k从1开始
	findKth = func(arr1 []int, arr2 []int, k int) int {
		if len(arr1) > len(arr2) {
			return findKth(arr2, arr1, k)
		}
		// 如果k值有越界的情况, 那么arr1就直接用完了
		if len(arr1) == 0 {
			return arr2[k-1]
		}
		// 为什么会有k==1的情况
		// 因为正常k是每次减少 k/2-1, 所以如果没有k越界的情况,一定会减到1
		if k == 1 {
			return min(arr1[0], arr2[0])
		}

		// 由于k是两个数字总长度的第k个,所以k有可能会大于单个数组的长度
		// 所以这里取下标时需要注意下,不要超出长度了
		// 由于k是第几个, 从1开始,所以这里ij下表需减去1
		i := min(len(arr1), k/2) - 1
		j := min(len(arr2), k/2) - 1

		if arr1[i] <= arr2[j] {
			return findKth(arr1[i+1:], arr2, k-i-1)
		} else {
			return findKth(arr1, arr2[j+1:], k-j-1)
		}
	}

	// 由于k从1开始, 因此中位数下标取值都是从1开始
	sumLen := len(nums1) + len(nums2)
	if sumLen%2 != 0 {
		return float64(findKth(nums1, nums2, sumLen/2+1))
	} else {
		return float64(findKth(nums1, nums2, sumLen/2)+findKth(nums1, nums2, sumLen/2+1)) / 2
	}

}
