package math

// https://leetcode.cn/problems/majority-element/
// 169. 多数元素
func majorityElement(nums []int) int {
	most := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == most {
			cnt++
		} else {
			if cnt == 0 {
				most = nums[i]
				cnt++
			} else {
				cnt--
			}
		}
	}
	return most
}
