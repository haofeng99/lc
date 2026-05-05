package hash

// https://leetcode.cn/problems/first-missing-positive/
// 41. 缺失的第一个正数

// 缺失的第一个正整数, 数量从1到len(nums)
// 1.如果为负数, 直接替换为不可能为结果的数据
// 2.将原属组中, 对应位置下标的数据取反以标记该数据出现过
// 3.遍历1-n, 如果对应下标位置的数据小于0, 则说明该数据出现过.
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 1. 不可能为结果的数替换
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1
		}
	}

	// 2. 对应位置取反来标记该数据出现过
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num > n {
			continue
		}
		if nums[num-1] < 0 {
			continue
		}
		nums[num-1] = -nums[num-1]
	}

	// 3. 遍历从1到n
	i := 1
	for ; i <= n; i++ {
		if nums[i-1] > 0 {
			return i
		}
	}
	return i
}

// 置换法
// 将对应数据放到具体下标位置上去
func firstMissingPositive_switch(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	i := 1
	for ; i <= n; i++ {
		if nums[i-1] != i {
			return i
		}
	}
	return i
}
