package sort

var ans493 int = 0

func reversePairs493(nums []int) int {
	splitNums(nums)
	return ans493
}

func splitNums(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	left := splitNums(nums[0:mid])
	right := splitNums(nums[mid:len(nums)])

	return mergeSortCount(left, right)
}

func mergeSortCount(left, right []int) []int {
	m, n := len(left), len(right)
	ans := make([]int, m+n)
	i, j, k := 0, 0, 0

	subCount := func(a1, a2 []int) int {
		p, ans := 0, 0
		for _, a := range a1 {
			for p < len(a2) && a > 2*a2[p] {
				p++
			}
			ans += p
		}
		return ans
	}

	ans493 += subCount(left, right)

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
	if j == n {
		for i < m {
			ans[k] = left[i]
			i++
			k++
		}
	}
	if i == m {
		for j < n {
			ans[k] = right[j]
			k++
			j++
		}
	}
	return ans
}

// https://leetcode.cn/problems/reverse-pairs/

func Show_reversePairs493() {

}
