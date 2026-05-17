package binarysearch

// 输出字典序最小的LIS
// 长度相同下, 数值越小越好
// https://www.nowcoder.com/questionTerminal/9cf027bf54714ad889d4f30ff0ae5481
// 1. 通过二分查找, 得到最终的最长的上升子序列的长度, 以及以每个位置作为结尾时的LIS的长度;
// 2. 从右往左构造答案：因为 tails 数组中相同位置 j 的值随时间单调不增，所以相同 dp 值的最右元素一定是最小的
func accessLIS(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}

	// dp[i] = 以 arr[i] 结尾的 LIS 长度
	dp := make([]int, n)
	tails := make([]int, 0, n)

	for i, x := range arr {
		if len(tails) == 0 || x > tails[len(tails)-1] {
			tails = append(tails, x)
			dp[i] = len(tails)
			continue
		}
		l, r := 0, len(tails)-1
		for l <= r {
			mid := l + (r-l)/2
			if tails[mid] >= x {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		tails[l] = x
		dp[i] = l + 1
	}

	maxLen := len(tails)
	res := make([]int, maxLen)
	target := maxLen

	// 只要保证从右往左遍历, 拿到的结果就是字典序最小的
	// tails[j] 的值随时间单调不增（patience sorting 每次只把 tails[j] 替换成更小的值）。因此对于相同 dp = k
	// 的所有元素，越靠右的出现得越晚，值一定越小。从右往左扫描并取第一个匹配，等价于在每个长度上取最小值，最终得到字典序最小的 LIS。
	for i := n - 1; i >= 0 && target > 0; i-- {
		if dp[i] == target {
			target--
			res[target] = arr[i]
		}
	}

	return res
}
