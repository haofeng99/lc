package binarysearch

// https://leetcode.cn/problems/search-a-2d-matrix/
// 74. 搜索二维矩阵
// 思路: 把整个二维数组拉长成为一维度数组, [每行第一个元素大于上一行最后一个元素的情形]
// O(logmn)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/n][mid%n]
		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
