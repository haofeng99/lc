package sort

// 327 区间和的个数
var ans327 int

func countRangeSum(nums []int, lower int, upper int) int {
	//rangeSumArr := rangeSum(nums)
	ans327 = 0
	return 0
}

func mergrSort327() {

}

func merge327(left, right []int) []int {
	m, n := len(left), len(right)
	i, j, k := 0, 0, 0
	ans := make([]int, m+n)
	for i < m && j < n {
		if left[i] <= right[j] {
			ans[k] = left[i]
			i++
		} else {
			ans[k] = right[j]
			j++
		}
		k++
	}
	if i == m {
		for j < n {
			ans[k] = right[j]
			j++
			k++
		}
	}
	if j == n {
		for i < m {
			ans[k] = left[i]
			i++
			k++

		}
	}
	return ans
}

func countRangeSumFromSortedArr(arr []int) int {
	return 0
}

func countRangeSumNum(left, right int) int {
	return (1 << (right - left + 1)) - 1
}

// 计算前缀和
func rangeSum(arr []int) []int {
	ans := make([]int, len(arr)+1)
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] + arr[i]
	}
	return ans
}

// https://leetcode.cn/problems/count-of-range-sum/description/
