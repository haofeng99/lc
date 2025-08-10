package sort

import "fmt"

func kClosest(points [][]int, k int) [][]int {
	quickSortK(points, 0, len(points)-1, k)
	return points[:k]
}

func quickSortK(points [][]int, left, right, k int) {

	if left >= right {
		return
	}
	l, r := left, right
	points[l], points[(l+r)/2] = points[(l+r)/2], points[l]
	pivot := calcLength(points[l])
	for l < r {
		for l < r && calcLength(points[r]) >= pivot {
			r--
		}
		for l < r && calcLength(points[l]) <= pivot {
			l++
		}

		if l < r {
			points[l], points[r] = points[r], points[l]
		}
	}
	if l > left {
		points[l], points[left] = points[left], points[l]
	}

	if l == k || l == k-1 {
		return
	}

	if l < k {
		quickSortK(points, l+1, right, k)
	} else {
		quickSortK(points, left, l-1, k)
	}
}

func calcLength(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

// https://leetcode.cn/problems/k-closest-points-to-origin/

func Show_KClosest() {
	points := [][]int{
		{3, 3},
		{-5, -1},
		{-2, 4},
	}
	k := 2
	res := kClosest(points, k)
	fmt.Println(res)
}
