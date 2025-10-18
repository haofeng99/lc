package dynamicprogramming

import "fmt"

// https://leetcode.cn/problems/maximum-subarray/
// 53. 最大子数组和

// 变种1: 输出对应子数组
// 变种2: https://www.nowcoder.com/discuss/820121. 剑指 Offer II 008

// 状态转移方程
// f(i)表示以第i个元素结尾的子数组的最高销售额
// f(i) = max(nums[i]+f(i-1), nums[i])
func maxSubArray(nums []int) int {

	maxSalesArr := make([]int, len(nums), len(nums))
	maxSalesArr[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		maxSalesArr[i] = max(nums[i], maxSalesArr[i-1]+nums[i])
	}

	ans := maxSalesArr[0]
	for i := 1; i < len(maxSalesArr); i++ {
		if maxSalesArr[i] > ans {
			ans = maxSalesArr[i]
		}
	}

	return ans
}

// 滚动数组优化
// 思路: maxSalesArr[i]只与maxSalesArr[i-1]和nums[i]有关系
// nums[i]在每一步都能取到, 因此只需要记录maxSalesArr[i-1]就行
// 直接把nums[i-1]存放maxSalesArr[i-1]的值
func maxSubArray_II(nums []int) int {

	// 为什么可以直接选取第一个元素, 因为题目要求非空, 最少需要选取一个元素
	// 默认只选取第一个元素时获取到最佳结果
	dp, ans := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])

		if dp > ans {
			ans = dp
		}
	}

	return ans
}

// todo
func maxSubArray_III(nums []int) int {

	dp, ans := nums[0], nums[0]
	left, right := 0, 0

	for i := 1; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])

		// 如果选取了当前数, 说明前面的数据都丢弃了, 数组从当前位开始
		if dp == nums[i] {
			left = i
		}

		// 每次有更新最终结果说明一定有新内容更新进来了
		if dp > ans {
			ans = dp
			right = i
		}
	}

	fmt.Println(nums[left : right+1])

	return ans
}
