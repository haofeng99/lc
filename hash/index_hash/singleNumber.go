package hash

// https://leetcode.cn/problems/single-number/
// 136. 只出现一次的数字

func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
