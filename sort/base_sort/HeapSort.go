package basesort

import "fmt"

/*
基础知识
对于一个下标为i的节点来说,它的左子节点下标位2*i+1;
最于一个下标为i的节点来说,它的右子节点下标是2*i+2;
对于一个下标为i的节点来说,它的父亲节点下标为(i-1)/2
*/
func HeapSort(arr []int) {
	length := len(arr)
	// 最后一个节点是下标是n-1,它的父节点是(idx-1)/2 -> (n-1-1)/2 = n/2-1
	for i := length/2 - 1; i >= 0; i-- {
		initMaxHeap(arr, i, length)
	}
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		initMaxHeap(arr, 0, i)
	}
}

// 初始化大顶堆
// i: 当前父节点下标
// len: 数组长度
func initMaxHeap(arr []int, i, len int) {
	p := i
	l, r := 2*i+1, 2*i+2
	if l < len && arr[l] > arr[p] {
		p = l
	}
	if r < len && arr[r] > arr[p] {
		p = r
	}
	if p != i {
		arr[p], arr[i] = arr[i], arr[p]
		initMaxHeap(arr, p, len)
	}
}

func ShowHeapSort() {
	arr := []int{5, 4, 3, 2, 1, 6, 7, 8, 10, 9}
	fmt.Printf("排序前: %v\n", arr)
	HeapSort(arr)
	fmt.Printf("排序后: %v\n", arr)
}
