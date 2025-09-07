package basesort

import "fmt"

/*
快速排序
*/
func quickSort(arr []int, left, right int) {
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
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)

}

/*
要实现逆序（降序）快速排序，只需将比较符号反过来即可：

arr[r] >= private 改为 arr[r] <= private
arr[l] <= private 改为 arr[l] >= private
*/
func quickSortDesc(arr []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	arr[l], arr[(l+r)/2] = arr[(l+r)/2], arr[l]
	private := arr[l]

	for l < r {
		for l < r && arr[r] <= private {
			r--
		}
		for l < r && arr[l] >= private {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	if l > left {
		arr[l], arr[left] = arr[left], arr[l]
	}
	quickSortDesc(arr, left, l-1)
	quickSortDesc(arr, l+1, right)
}

// 快排通常用于解决topK问题, 不关心K个数的顺序,只关心谁在K内

func ShowQuickSort() {
	arr := []int{3, 9, 0, 5, 6, 7, 1, 2, 4}
	fmt.Printf("排序前: %v\n", arr)
	quickSortDesc(arr, 0, len(arr)-1)
	fmt.Printf("排序后: %v\n", arr)
}
