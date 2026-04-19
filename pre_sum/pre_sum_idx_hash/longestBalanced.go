package presum

import "strings"

// 3900 一次交换后的最长平衡子串
// https://leetcode.cn/problems/longest-balanced-substring-after-one-swap/

// 思路: 525变种题
// 关键点:
// 1. 为什么要存储一个前缀和的最小的两个位置
// 2. 什么时候可以交换

// 关键信息(右边j, 左边i)
// 1. 交换后平衡, 即交换前子串内部0和1的数量差为2
// 2. 子串的长度为j-i, 则如果可以满足交换

// 交换右边的0 需要记录preSum-2前缀和出现的第一个位置
// 2.1 交换1出去, 此时1比0多2个 0的个数(j-i-2)/2 < zeroCnt
// 2.2 交换0出去, 此时0比1多2个 1的个数(j-i-2)/2 < oneCnt

// 向左交换 需要记录preSum-2前缀和出现的第二个位置(如果有的话)
// 如果j-i-2/2 == zeroCnt, 说明子串内已经包含全部的0, 右边无0可以借用
// 如果有同一个前缀和的两个位置, 则说明两个位置之间存在相同数量的0和1, 那就取j到第二个i的长度
func longestBalanced(s string) int {
	n := len(s)

	zeroCnt := strings.Count(s, "0")
	oneCnt := n - zeroCnt

	ans := 0
	pre_sum := 0

	mp := map[int][]int{0: {-1}}
	for i := 0; i < n; i++ {
		if s[i] == byte('0') {
			pre_sum--
		} else {
			pre_sum++
		}

		p := mp[pre_sum]
		if len(p) < 2 {
			p = append(p, i)
			mp[pre_sum] = p
		}

		// 0 不做交换
		if pre, ok := mp[pre_sum]; ok {
			if len(pre) == 2 {
				ans = max(ans, i-pre[0])
			}
		}

		// 2 交换外部的0
		if pre, ok := mp[pre_sum-2]; ok {
			if (i-pre[0]-2)/2 < zeroCnt {
				ans = max(ans, i-pre[0])
			} else if len(pre) == 2 {
				ans = max(ans, i-pre[1])
			}
		}

		// -2 交换外部的1
		if pre, ok := mp[pre_sum+2]; ok {
			if (i-pre[0]-2)/2 < oneCnt {
				ans = max(ans, i-pre[0])
			} else if len(pre) == 2 {
				ans = max(ans, i-pre[1])
			}
		}

	}
	return ans
}
