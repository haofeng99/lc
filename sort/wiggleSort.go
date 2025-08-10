package sort

import "fmt"

// 不满足条件,直接和前面的数据交换,已经遍历过的数据依旧满足题目要求条件,且更满足
func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if (i%2 == 1 && nums[i] < nums[i-1]) || (i%2 == 0 && nums[i] > nums[i-1]) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

// https://www.lintcode.com/problem/508/

func Show_wiggleSort() {
	nums := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}
