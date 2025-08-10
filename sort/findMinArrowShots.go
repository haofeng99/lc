package sort

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 1
	curRight := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= curRight {
			curRight = min(curRight, points[i][1])
			continue
		} else {
			curRight = points[i][1]
			ans++
		}
	}

	return ans
}

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
func Show_findMinArrowShots() {

}
