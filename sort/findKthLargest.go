package sort

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
// 主要考察对堆排序和快速排序的理解
// 215 数组中的第K个最大元素

func findKthLargest(nums []int, k int) int {
	var quickSort func([]int, int, int) int
	quickSort = func(arr []int, left, right int) int {
		if left >= right {
			return arr[len(arr)-k]
		}
		l, r := left, right
		arr[l], arr[(l+r)/2] = arr[(l+r)/2], arr[l]
		p := arr[l]
		for l < r {
			for l < r && arr[r] >= p {
				r--
			}
			for l < r && arr[l] <= p {
				l++
			}
			if l != r {
				arr[l], arr[r] = arr[r], arr[l]
			}
		}
		if l > left {
			arr[l], arr[left] = arr[left], arr[l]
		}
		if r == len(arr)-k {
			return arr[r]
		} else if r > len(arr)-k {
			return quickSort(arr, left, l-1)
		} else {
			return quickSort(arr, r+1, right)
		}
	}
	return quickSort(nums, 0, len(nums)-1)
}
