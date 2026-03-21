package hash

// https://leetcode.cn/problems/find-the-duplicate-number/
// 287. 寻找重复数
func findDuplicate(nums []int) int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] < 0 {
			return num
		}
		nums[num-1] = -nums[num-1]
	}
	return -1
}
