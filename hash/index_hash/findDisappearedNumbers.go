package hash

// https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
// 448. 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	ans := []int{}

	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
