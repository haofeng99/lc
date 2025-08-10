package sort

import "fmt"

func pancakeSort(arr []int) []int {
	ans := []int{}
	right := len(arr) - 1
	for right > 0 {
		maxIdx := 0
		for i := 0; i <= right; i++ {
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}
		if right != maxIdx {
			reserve(arr, 0, maxIdx)
			ans = append(ans, maxIdx+1)
			reserve(arr, 0, right)
			ans = append(ans, right+1)
		}
		right--
	}
	return ans
}

func reserve(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// https://leetcode.cn/problems/pancake-sorting/
// 题解: https://leetcode.cn/problems/pancake-sorting/solutions/1275857/fu-xue-ming-zhu-tu-jie-jian-bing-pai-xu-jto8z/
// 大致思路: 先找局部最大值, 放到首位,再放到局部尾,最后不断缩小局部范围

func Show_pancakeSort() {
	arr := []int{1, 4, 2, 3}
	ans := pancakeSort(arr)
	fmt.Println(ans)

	// 2  4123
	// 4  3214
	// 1  3214
	// 3  1234
	// 1  1234
}
