package dynamicprogramming

import "fmt"

// 152. 乘积最大子数组
// 腾讯 字节 虾皮 华为

// 解题思路:
// 1.它和最大子数组和不同, 不能直接在每一步求最大值, 因为负数乘负数可以得到正数
// 2.因此我们在每一步需要同时求最大值和最小值, 并将最大值和最小值与当前位相乘, 多个数求最大
// 3.状态转移方程: f_max = max(f_max_pre * cur, f_min_pre * cur, cur),其中f_min = min(f_max_pre * cur, f_min_pre * cur, cur)
func maxProduct(nums []int) int {
	// 为什么初始值可以是1:因为任何数乘以1保持不变
	// 且题目要求最大非空, 因此最少选取一个元素, 因此可以以1作为初始值
	f_max, f_min := nums[0], nums[0]
	ans := f_max // 默认只选取第一个数时得到最优结果
	for i := 1; i < len(nums); i++ {
		// f_max和f_min必须同时计算, 不能先计算一个后再计算另一个
		f_max, f_min = max(f_max*nums[i], f_min*nums[i], nums[i]),
			min(f_max*nums[i], f_min*nums[i], nums[i])
		if f_max > ans {
			ans = f_max
		}
	}

	return ans
}

// 进阶: 输出该子数组
func maxProduct_subarray(nums []int) int {
	f_max, f_min := nums[0], nums[0]
	ans := f_max
	right := 0
	for i := 1; i < len(nums); i++ {
		f_max, f_min = max(f_max*nums[i], f_min*nums[i], nums[i]),
			min(f_max*nums[i], f_min*nums[i], nums[i])

		if f_max > ans {
			ans = f_max
			right = i
		}
	}

	left := right
	tmpAns := ans
	if tmpAns != 0 {
		for tmpAns != 1 {
			tmpAns = tmpAns / nums[left]
			left--
		}
		left++
	}

	fmt.Println(nums[left : right+1])

	return ans
}
