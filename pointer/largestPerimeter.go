package pointer

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

// https://leetcode.cn/problems/largest-perimeter-triangle/

// 逆向枚举,为什么只枚举三个相邻的数据

// 如果 A[i-2] + A[i-1] <= A[i] ，这三个数一定不能构成三角形，而A[i-3]以及更往前的数，都小于等于A[i-2]，
// 所以再往前取任何两个数只会让相加的值更小，就更不能满足 A[j] + A[k] > A[i]了 (j<i-2, k<i-1, j<k)。
// 所以如果相邻的数构不成三角形，就不需要再固定第三个数并往前找两个数了

// 如果 A[i-2] + A[i-1] > A[i]j，这三个数可以构成三角形，再往前找只会让周长变短，所以也不用再往前了。
