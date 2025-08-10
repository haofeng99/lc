package sort

func carFleet(target int, position []int, speed []int) int {
	ans, n := 0, len(position)-1
	costTm := make([]float32, target)
	for i := 0; i < n; i++ {
		costTm[position[i]] = float32(target-position[i]) / float32(speed[i])
	}

	// 拦车
	maxCost := float32(0)
	for i := target - 1; i >= 0; i-- {
		if costTm[i] <= maxCost {
			continue
		}
		ans++
		maxCost = costTm[i]
	}
	return ans
}

// https://leetcode.cn/problems/car-fleet/
