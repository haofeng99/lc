package onedim

// https://leetcode.cn/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	preSumMap := make(map[int]int)
	ans := 0
	curPreSum := 0
	for i := 0; i < len(nums); i++ {
		curPreSum += nums[i]

		if curPreSum == k {
			ans++
		}

		targetPreSum := curPreSum - k

		if _, ok := preSumMap[targetPreSum]; ok {
			ans += preSumMap[targetPreSum]
		}
		preSumMap[curPreSum] += 1
	}

	return ans
}
