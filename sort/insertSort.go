package sort

import "fmt"

// 插入排序 O(n^2) O(1)
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func ShowInsertSort() {
	arr := []int{4, 9, 8, 3, 4, 5, 1, 2}
	fmt.Printf("排序前: %v\n", arr)
	insertSort(arr)
	fmt.Printf("排序后: %v\n", arr)
}
