package bit

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		if calcNum1Count(a) == calcNum1Count(b) {
			return a < b
		}
		return calcNum1Count(a) < calcNum1Count(b)
	})
	return arr
}

// 计算一个数二进制下1的数量
func calcNum1Count(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 2
		num /= 2
	}
	return ans
}

var bit = [1e4 + 1]int{}

func calcNum1Count2() {
	for i := 1; i <= 1e4; i++ {
		bit[i] = bit[i>>1] + i&1
	}
}
