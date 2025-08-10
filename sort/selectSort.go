package sort

import "fmt"

// 选择排序 O(n^2) O(1) 不稳定
func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx, minValue := i, arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minValue {
				minIdx, minValue = j, arr[j]
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

func ShowSelectSort() {
	arr := []int{7, 8, 2, 4, 6, 1, 3, 5, 8}
	fmt.Printf("排序前: %v\n", arr)
	selectSort(arr)
	fmt.Printf("排序后: %v\n", arr)
}
