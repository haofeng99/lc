package bfs

import "fmt"

// 从(0,0)出发，相邻元素为子节点，BFS结果即为所求。注意要优先遍历下方的元素，再遍历右方的元素。
func findDiagonalOrder(nums [][]int) []int {
	ans := make([]int, 0)

	see := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		see[i] = make([]int, len(nums[i]))
	}

	queue := [][]int{{0, 0}}
	for len(queue) > 0 {
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			pos := queue[i]
			i, j := pos[0], pos[1]
			ans = append(ans, nums[i][j])
			if i+1 < len(nums) && j < len(nums[i+1]) && see[i+1][j] == 0 {
				queue = append(queue, []int{i + 1, j})
				see[i+1][j] = 1
			}
			if i < len(nums) && j+1 < len(nums[i]) && see[i][j+1] == 0 {
				queue = append(queue, []int{i, j + 1})
				see[i][j+1] = 1
			}
		}
		queue = queue[curLen:]
	}

	return ans
}

// https://leetcode.cn/problems/diagonal-traverse-ii/

func Show_findDiagonalOrder() {
	nums := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}
	ans := findDiagonalOrder(nums)
	fmt.Println(ans)
}
