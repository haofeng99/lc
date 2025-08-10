package sort

import "fmt"

// 横纵坐标之间没联系,可分开计算
func minTotalDistance(grid [][]int) int {
	iPos := make([]int, 0)
	jPos := make([]int, 0)
	for i, row := range grid {
		for _, col := range row {
			if col == 1 {
				iPos = append(iPos, i)
			}
		}
	}
	for j, _ := range grid[0] {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 1 {
				jPos = append(jPos, j)
			}
		}
	}
	ans := 0
	for l, r := 0, len(iPos)-1; l < r; l, r = l+1, r-1 {
		ans += (iPos[r] - iPos[l])
	}
	for l, r := 0, len(jPos)-1; l < r; l, r = l+1, r-1 {
		ans += (jPos[r] - jPos[l])
	}
	return ans
}

// https://cloud.tencent.com/developer/article/1787757
// https://www.lintcode.com/problem/912/

func Show_minTotalDistance() {
	grid := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println(minTotalDistance(grid))
}
