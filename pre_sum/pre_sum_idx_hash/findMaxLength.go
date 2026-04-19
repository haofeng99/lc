package presum

// 525. 连续子数组
// https://leetcode.cn/problems/contiguous-array/

// 思路1:
// map k: 前缀和值. v: 前缀和值出现的下标位置
// 如果前缀和为0, 则取最远的下标位置
// 如果前缀和不为0, 则取下表位置的最差值
func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	preSumIdxsMap := make(map[int][]int)

	pre_sum := make([]int, len(nums))
	pre_sum[0] = nums[0]
	preSumIdxsMap[pre_sum[0]] = append(preSumIdxsMap[pre_sum[0]], 0)
	for i := 1; i < len(nums); i++ {
		pre_sum[i] = pre_sum[i-1] + nums[i]
		preSumIdxsMap[pre_sum[i]] = append(preSumIdxsMap[pre_sum[i]], i)
	}

	ans := 0
	for k, v := range preSumIdxsMap {
		if k == 0 {
			ans = max(ans, v[len(v)-1]+1)
		} else {
			ans = max(ans, v[len(v)-1]-v[0])
		}
	}

	return ans
}

// 优化过程
// 1. 遍历原数组过程中前缀和就可以求出
// 2. 前缀和也不需要定义数组, 直接一个变量就可以实现
// 3. map k为前缀和 v为该前缀和第一次出现的位置
// 4. 针对前缀和为0的场景特殊处理 可以先把{0, -1}初始化到map中
// 5. 注意点: 子串的长度为i-j (考虑前缀和的计算过程)
func findMaxLength_ii(nums []int) int {
	n := len(nums)
	pre_sum := 0
	pre_sum_map := map[int]int{0: -1}

	// 因为涉及到改动原数组的行为, 所以前缀和需要从头开始算
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		pre_sum += nums[i]

		if pre_idx, contions := pre_sum_map[pre_sum]; contions {
			ans = max(ans, i-pre_idx)
		} else {
			pre_sum_map[pre_sum] = i
		}
	}

	return ans
}
