package pointer

import "fmt"

func sortArrayByParityII(nums []int) []int {
	j, o := 0, 1
	for j < len(nums) && o < len(nums) {
		for j < len(nums) && suitPos(j, nums[j]) {
			j += 2
		}
		for o < len(nums) && suitPos(o, nums[o]) {
			o += 2
		}
		if j < len(nums) && o < len(nums) {
			nums[j], nums[o] = nums[o], nums[j]
		}
	}
	return nums
}

func suitPos(idx, num int) bool {
	return ((idx%2 == 0 && num%2 == 0) || (idx%2 != 0 && num%2 != 0))
}

// https://leetcode.cn/problems/sort-array-by-parity-ii/

func Show_sortArrayByParityII() {
	nums := []int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3}
	res := sortArrayByParityII(nums)
	fmt.Println(res)
}
