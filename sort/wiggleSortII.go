package sort

import (
	"fmt"
	"sort"
)

// 一.普通排序
// 1. 如果所有输入数组都可以满足要求,则数组中相同元素的数量不可能超过(n + 1) / 2
// 2. 数组进行拆分时, 如果数组长度为奇数, 则前面的数组长度一般要多一个元素
// 3. 排序后, 分组, 分组后结果逆序, 最后组合最后结果
func wiggleSortII(nums []int) {
	arr := append([]int{}, nums...)
	sort.Ints(arr)
	l, r, i := (len(arr)+1)/2-1, len(arr)-1, 0
	for l >= 0 {
		if l >= 0 {
			nums[i] = arr[l]
			i++
		}
		if r >= (len(arr)+1)/2 {
			nums[i] = arr[r]
			i++
		}
		l--
		r--
	}
}

// O(n)
func wiggleSortII_I(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 找到中位数
	arr := append([]int{}, nums...)
	mid := quickSelect(arr, 0, len(arr)-1, (n-1)/2)

	// three-way-partition
	three_way_partition(arr, mid)

	// 前数组元素个数 >= 后数组元素个数
	l, r, i := (n+1)/2-1, n-1, 0
	for l >= 0 {
		if l >= 0 {
			nums[i] = arr[l]
			i++
		}
		if r >= (n+1)/2 {
			nums[i] = arr[r]
			i++
		}
		l--
		r--
	}
}

// three-way-partition
// p : 基准值, 根据该基准值进行排序, 排序后, 所有小于p的值都在左边,大于p的值都在右边
func three_way_partition(arr []int, p int) {
	i, j, n := 0, 0, len(arr)-1
	for j <= n {
		if arr[j] > p {
			arr[n], arr[j] = arr[j], arr[n]
			n--
		} else if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			j++
			i++
		} else {
			j++
		}
	}
}

// 返回数据nums中,第k个大的元素
func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	nums[r], nums[(l+r)/2] = nums[(l+r)/2], nums[r]
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	}
	if p < k {
		return quickSelect(nums, p+1, r, k)
	} else {
		return quickSelect(nums, l, p-1, k)
	}
}

// 单独分区过程-小于基准值的放左边,大于基准值的放右边
func partition(nums []int, l, r int) int {
	pivot := nums[r]
	idx := l // idx 左边的值都严格小于pivot
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
	nums[r], nums[idx] = nums[idx], nums[r]
	return idx
}

// 三、桶排序

// https://leetcode.cn/problems/wiggle-sort-ii/description/
// 题解: https://leetcode.cn/problems/wiggle-sort-ii/solutions/45144/yi-bu-yi-bu-jiang-shi-jian-fu-za-du-cong-onlognjia/
func Show_wiggleSortII() {
	nums := []int{1, 2, 1, 2, 1, 1, 2, 2, 1}
	wiggleSortII_I(nums)
	// three_way_partition(nums, 4)
	fmt.Println(nums)
}

func Show_partition() {
	nums := []int{3, 4, 5, 7, 8, 9, 1, 2, 3, 4, 6, 7}
	idx := partition(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println(idx)
}
