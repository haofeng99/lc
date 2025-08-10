package sort

import (
	"fmt"
	"testing"
)

func _sortArray(nums []int) []int {
	return _mergeSort(nums)
}

// i left:2*i+1 right:2*i+2
// i parent:(i-1)/2
// n-1 parent:(n-1-1)/2 -> n/2-1
func _heapSort(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		_initMaxHeap(nums, i, n)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		_initMaxHeap(nums, 0, i)
	}
}

func _initMaxHeap(nums []int, i, length int) {
	f := i
	l, r := 2*i+1, 2*i+2
	if l < length && nums[l] > nums[f] {
		f = l
	}
	if r < length && nums[r] > nums[f] {
		f = r
	}
	if f != i {
		nums[f], nums[i] = nums[i], nums[f]
		_initMaxHeap(nums, f, length)
	}
}

func _mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) >> 1
	left := _mergeSort(nums[:mid])
	right := _mergeSort(nums[mid:])
	return _merge(left, right)
}

func _merge(left, right []int) []int {
	m, n := len(left), len(right)
	ans := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if left[i] <= right[j] {
			ans[k] = left[i]
			i++
			k++
		} else {
			ans[k] = right[j]
			j++
			k++
		}
	}
	for i < m {
		ans[k] = left[i]
		i++
		k++
	}
	for j < n {
		ans[k] = right[j]
		j++
		k++
	}
	return ans
}

func _quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	mid := l + (r-l)>>1
	flagValue := nums[mid]
	nums[l], nums[mid] = nums[mid], nums[l]
	for l < r {
		for l < r && nums[r] >= flagValue {
			r--
		}
		for l < r && nums[l] <= flagValue {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	if l > left {
		nums[l], nums[left] = nums[left], nums[l]
	}
	_quickSort(nums, left, l-1)
	_quickSort(nums, l+1, right)
}

func TestSort(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	arr := []int{3, 2, 1, 4, 5, 6, 7, 8, 9, 10}
	got := _mergeSort(arr)
	fmt.Println(got)
}

func TestMakeFunc(t *testing.T) {
	arr := make([]int, 10)
	fmt.Println(arr)
}

func TestMerge(t *testing.T) {
	left := []int{1, 3, 5, 7, 9}
	right := []int{2, 4, 6, 8, 10}
	got := _merge(left, right)
	fmt.Println(got)
}
