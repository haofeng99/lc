package basesort

import "fmt"

func ShellSort(arr []int) {
	for s := len(arr) / 2; s > 0; s /= 2 {
		for i := s; i > 0; i -= s {
			for j := i; j > 0; i -= s {
				if arr[i] < arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
}

func ShowShellSort() {
	arr := []int{7, 9, 10, 12, 6, 3, 1}
	fmt.Printf("排序前: %v\n", arr)
	ShellSort(arr)
	fmt.Printf("排序后: %v\n", arr)
}
