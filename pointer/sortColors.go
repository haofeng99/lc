package pointer

import "fmt"

// [0, p0] = 0
// (p0, i) = 1
// (p2, len - 1] == 2
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i < p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

// three-way-partition
func sortColors3(nums []int) {
	i, j, n := 0, 0, len(nums)-1
	p := 1
	for j <= n {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] > p {
			nums[j], nums[n] = nums[n], nums[j]
			n--
		} else {
			j++
		}
	}
}

// 刷房子
func sortColors2(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		nums[i] = 2
		if num < 2 {
			nums[p1] = 1
			p1++
		}
		if num < 1 {
			nums[p0] = 0
			p0++
		}
	}
}

// https://leetcode.cn/problems/sort-colors/description/

func Show_sortColors() {
	nums := []int{2, 1, 2, 0, 2, 1, 0, 1, 0, 1}
	sortColors2(nums)
	fmt.Printf("分类后:%v\n", nums)
}
