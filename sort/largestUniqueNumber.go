package sort

// 1133. 最大唯一数
// 给你一个整数数组 A，请找出并返回在该数组中仅出现一次的最大整数。
// 如果不存在这个只出现一次的整数，则返回 -1。
// 备注: 1 <= A.length <= 2000.   0 <= A[i] <= 1000
func largestUniqueNumber(A []int) int {
	cnt := [1001]int{} // cnt类型是数组而不是slice
	for _, num := range A {
		cnt[num]++
	}
	for i := 1000; i >= 0; i-- {
		if cnt[i] == 1 {
			return i
		}
	}
	return -1
}

func TestlargestUniqueNumber(arr []int) int {
	return largestUniqueNumber(arr)
}
