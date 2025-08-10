package sort

import "fmt"

func smallestK(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSort(arr []int, left, right, k int) {
	if left >= right {
		return
	}
	l, r := left, right
	arr[l], arr[(l+r)/2] = arr[(l+r)/2], arr[l]
	private := arr[l]

	for l < r {
		for l < r && arr[r] >= private {
			r--
		}
		for l < r && arr[l] <= private {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	if l > left {
		arr[l], arr[left] = arr[left], arr[l]
	}
	if l == k || l == k-1 {
		return
	}
	if l > k {
		quickSort(arr, left, l-1, k)
	} else {
		quickSort(arr, l+1, right, k)
	}
}

// https://leetcode.cn/problems/smallest-k-lcci/
// 右边也需要排序

func Show_smallestK() {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	k := 4
	res := smallestK(arr, k)
	fmt.Println(res)
}
