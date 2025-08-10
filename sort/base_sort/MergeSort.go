package basesort

import "fmt"

func MergeSort(arr []int) []int {
	length := len(arr)

	if length < 2 {
		return arr
	}

	mid := length / 2

	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:length])

	return merge(left, right)
}

func merge(left, right []int) []int {
	m, n := len(left), len(right)
	res := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if left[i] <= right[j] {
			res[k] = left[i]
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
			i = i + 1
			k = k + 1
		}
	}
	return res
}

// 归并排序通常用于求解逆序对问题

func ShowMergeSort() {
	arr := []int{3, 9, 0, 5, 6, 7, 1, 2, 4}
	fmt.Printf("排序前: %v\n", arr)
	res := MergeSort(arr)
	fmt.Printf("排序后: %v\n", res)
}
