package dynamicprogramming

// 2708 一个小组的最大实力值

// 思路:
// 1. 152题的子数组场景时, 必须选当前元素, 而子小组可以不选当前元素,因此状态转移方程相较于152需要增加当前元素不选取的情况
// 2. f_max = max(f_max_pre * cur, f_min_pre * cur, cur, f_max_pre),其中f_min = min(f_max_pre * cur, f_min_pre * cur, cur, f_min_pre)
func maxStrength(nums []int) int64 {
	f_max, f_min := nums[0], nums[0]
	ans := f_max
	for i := 1; i < len(nums); i++ {
		f_max, f_min = max(f_max*nums[i], f_min*nums[i], nums[i], f_max),
			min(f_max*nums[i], f_min*nums[i], nums[i], f_min)
		if f_max > ans {
			ans = f_max
		}
	}

	return int64(ans)
}
