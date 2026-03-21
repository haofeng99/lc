package hash

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/
// 442. 数组中重复的数据

// 思路: 原地hash
func findDuplicates(nums []int) []int {
	ans := []int{}
	for _, num := range nums {
		if num < 0 {
			num = num * -1
		}
		if nums[num-1] < 0 {
			ans = append(ans, num)
			continue
		}
		nums[num-1] = -1 * nums[num-1]
	}
	return ans
}
