package basesort

import "fmt"

// 冒泡排序 O(n^2) O(1) 稳定
func bubbleSort(arr []int) {
	for t := 0; t < len(arr)-1; t++ {
		needSwap := false
		for i := 1; i < len(arr)-t; i = i + 1 {
			if arr[i] < arr[i-1] {
				needSwap = true
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		if !needSwap {
			return
		}
	}
}

func ShowBubbleSort() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Printf("排序前: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("排序后: %v\n", arr)
}
