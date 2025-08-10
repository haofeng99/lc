package sort

import "sort"

func subSort(array []int) []int {
	if sort.IntsAreSorted(array) {
		return []int{-1, -1}
	}
	copyArray := append([]int{}, array...)
	sort.Ints(copyArray)
	left := 0
	for left < len(array) {
		if array[left] != copyArray[left] {
			break
		}
		left++
	}
	right := len(array) - 1
	for right >= 0 {
		if array[right] != copyArray[right] {
			break
		}
		right--
	}
	if left != right && left < right {
		return []int{left, right}
	}
	return []int{-1, -1}

}

func subSort2(array []int) []int {
	l, r := -1, -1
	if len(array) == 0 {
		return []int{l, r}
	}
	min, max := array[len(array)-1], array[0]

	for i := 0; i < len(array); i++ {
		if array[i] < max {
			r = i
		} else {
			max = array[i]
		}

		j := len(array) - 1 - i
		if array[j] > min {
			l = j
		} else {
			min = array[j]
		}
	}
	return []int{l, r}
}

// https://leetcode.cn/problems/sub-sort-lcci/
