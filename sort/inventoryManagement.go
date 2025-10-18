package sort

// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
// LCR 159. 库存管理 III
// 思路: 最小k问题 快排 + 堆排序
func inventoryManagement(stock []int, cnt int) []int {
	var quickSort func([]int, int, int) []int
	quickSort = func(arr []int, left, right int) []int {
		if left >= right {
			return stock[:cnt]
		}
		l, r := left, right
		p := l + (r-l)/2
		arr[l], arr[p] = arr[p], arr[l]
		pv := arr[l]
		for l < r {
			for l < r && pv <= arr[r] {
				r--
			}
			for l < r && pv >= arr[l] {
				l++
			}
			if l < r {
				arr[l], arr[r] = arr[r], arr[l]
			}
		}
		if l > left {
			arr[left], arr[l] = arr[l], arr[left]
		}
		if l+1 == cnt || l == cnt {
			return stock[:cnt]
		} else {
			if l < cnt {
				return quickSort(arr, r+1, right)
			} else {
				return quickSort(arr, left, l-1)
			}
		}
	}
	return quickSort(stock, 0, len(stock)-1)
}

func inventoryManagement_heapSort(stock []int, cnt int) []int {
	// 下标为i的坐子节点 2*i+1, 右子节点2*i+2
	// 下标为i的节点的父节点为 (i-1)/2
	// 最后一个下标为n-1, 则最后一个父节点为 (n-1-1)/2 -> n/2-1

	var buildMaxHeap func([]int, int, int)
	buildMaxHeap = func(arr []int, i, len int) {
		p := i
		l, r := 2*i+1, 2*i+2
		if l < len && arr[l] > arr[p] {
			p = l
		}
		for r < len && arr[r] > arr[p] {
			p = r
		}
		if p != i {
			arr[p], arr[i] = arr[i], arr[p]
			buildMaxHeap(arr, p, len)
		}
	}

	n := len(stock)
	for i := n/2 - 1; i >= 0; i-- {
		buildMaxHeap(stock, i, n)
	}

	for i := n - 1; i > cnt-1; i-- {
		stock[0], stock[i] = stock[i], stock[0]
		buildMaxHeap(stock, 0, i)
	}

	return stock[:cnt]
}
