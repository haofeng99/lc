package sort

import (
	"math"
)

// https://leetcode.cn/problems/maximum-gap/description/
// 简单的计数排序无法满足要求,因为一些测试用例友好, 例如[2,99999999],cannot allocate memory
func maximumGap(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	count := make([]int, max+1)
	for _, v := range nums {
		count[v]++
	}
	ans, pre := 0, -1
	for i, v := range count {
		if v != 0 {
			if pre != -1 {
				ans = maxInt(ans, i-pre)
			}
			pre = i
		}
	}
	return ans
}

// 进阶结论
// 元素之间的最大间距一定不会出现在某个桶的内部，而一定会出现在不同桶当中
type bucketMaxAndMin struct {
	max int // 桶内最大值
	min int // 桶内最小值
}

func newBucketMaxAndMin(maxValue, minValue int) *bucketMaxAndMin {
	return &bucketMaxAndMin{
		max: minValue,
		min: maxValue,
	}
}

func _bucketSortHelper(arr []int) int {
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	bucketNum := int(math.Sqrt(float64(len(arr))))
	bucket := make([][]int, bucketNum)
	for j := 0; j < len(arr); j++ {
		// 分桶公式
		n := bucketNum * (arr[j] - min) / (max - min + 1)
		bucket[n] = append(bucket[n], arr[j])
		// 桶内单轮次排序
		k := len(bucket[n]) - 2
		for k >= 0 && bucket[n][k+1] < bucket[n][k] {
			bucket[n][k], bucket[n][k+1] = bucket[n][k+1], bucket[n][k]
			k--
		}
	}

	ans, pre := 0, -1
	for i, v := range bucket {
		bucketLen := len(v)
		if bucketLen != 0 {
			if pre != -1 {
				ans = maxInt(ans, v[0]-bucket[pre][len(bucket[pre])-1])
			}
			ans = maxInt(ans, maxGap(v))
			pre = i
		}
	}
	return ans
}

func maxGap(nums []int) int {
	ans := 0
	if len(nums) < 2 {
		return ans
	}
	pre := -1
	for i, v := range nums {
		if pre != -1 {
			ans = maxInt(ans, v-nums[pre])
		}
		pre = i
	}
	return ans
}
