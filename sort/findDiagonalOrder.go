package sort

import "fmt"

// 排序
// 1.优先按照下标和排序正向,其次按照第一个坐标逆向排,最后按横坐标正向排序
func findDiagonalOrder(nums [][]int) []int {
	// TODO
	return nil
}

// 遍历, 按下标和相等的数据聚合
// 主对角线的数据, i+j相等
// 副对角线的数据, i-j相等
func findDiagonalOrder2(nums [][]int) []int {
	path := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i+j < len(path) {
				path[i+j] = append(path[i+j], nums[i][j])
			} else {
				path = append(path, []int{nums[i][j]})
			}
		}
	}
	ans := []int{}
	for i := 0; i < len(path); i++ {
		for j := len(path[i]) - 1; j >= 0; j-- {
			ans = append(ans, path[i][j])
		}
	}
	return ans
}

// https://leetcode.cn/problems/diagonal-traverse-ii/
func Show_findDiagonalOrder() {
	nums := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}
	ans := findDiagonalOrder2(nums)
	fmt.Println(ans)
}
