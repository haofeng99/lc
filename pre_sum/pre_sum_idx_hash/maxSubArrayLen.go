package presum

// https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/description/
// https://www.lintcode.com/problem/911/description
// 最大子数组之和为k
/**
 * @param nums: an array
 * @param k: a target value
 * @return: the maximum length of a subarray that sums to k
 */
func MaxSubArrayLen(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := 0
	preSum := 0
	mp := map[int]int{0: -1} // 不能遗忘

	for i, num := range nums {
		preSum += num
		leftPreSum := preSum - k

		if preI, ok := mp[leftPreSum]; ok {
			ans = max(ans, i-preI)
		}

		if _, ok := mp[preSum]; !ok {
			mp[preSum] = i
		}

	}
	return ans
}
