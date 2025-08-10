package basesort

func CountSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}
	idx := 0
	for i, v := range count {
		for j := 0; j < v; j++ {
			arr[idx] = i
			idx++
		}
	}
	return arr
}
