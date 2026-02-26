package presum

import (
	"math"
)

// 53. 最大子数组和
// 1. 计算原数组的前缀和数组
// 2. 在前缀和数组中, 计算某两个下标之间的最大差
// 3. 将最大差和数组和进行比对
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	preSumArr := make([]int, len(nums)+1)
	for i, v := range nums {
		preSumArr[i+1] = preSumArr[i] + v
	}

	_, _, maxSubArrSum := maxDifference(preSumArr)

	return maxSubArrSum
}

// 前缀和之后本质就是求数组中的最大差值
// 121 买卖股票的最佳时机
// 找到最大差值 arr[j] - arr[i] (j > i)
func maxDifference(arr []int) (int, int, int) {
	n := len(arr)

	minIndex := 0 // 到目前为止最小值的下标
	maxDiff := math.MinInt32
	buyIndex, sellIndex := -1, -1

	for i := 1; i < n; i++ {
		// 计算如果今天卖出的利润
		diff := arr[i] - arr[minIndex]

		// 更新最大差值
		if diff > maxDiff {
			maxDiff = diff
			buyIndex = minIndex
			sellIndex = i
		}

		// 更新最小值下标
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	return buyIndex, sellIndex, maxDiff
}
