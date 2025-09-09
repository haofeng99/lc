package search

// https://leetcode.cn/problems/search-a-2d-matrix-ii/
// 240. 搜索二维矩阵 II

// 思路: 从右上角往左下角搜索
// O(m+n) O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 变种1
// 矩阵有重复值，要找到横纵坐标和最小的目标值
// 思路: 找到一个元素以后, 不直接返回, 继续向左找
func searchMatrixWithDupValue(matrix [][]int, target int) (int, int) {
	m, n := len(matrix), len(matrix[0])
	resI, resJ := m, n
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			if (i + j) < (resI + resJ) {
				resI, resJ = i, j
			}
			j-- // 继续向左找
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return resI, resJ
}

// 变种2
// 时间复杂度为O(logn) 或小于O(m+n)
// 对角线分治+二分查找，严格优于O(m+n)
func searchMatrixBestTimeComplex(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	var search func(row, col, height, width int) bool
	search = func(row, col, height, width int) bool {
		if height <= 0 || width <= 0 {
			return false
		}
		minDim := height
		if width < height {
			minDim = width
		}
		// 在对角线方向二分
		lo, hi := 0, minDim-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			v := matrix[row+mid][col+mid]
			if v == target {
				return true
			} else if v < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		// 分治：排除对角线左上和右下的区域
		return search(row+lo, col, height-lo, lo) || search(row, col+lo, lo, width-lo)
	}
	return search(0, 0, m, n)
}
