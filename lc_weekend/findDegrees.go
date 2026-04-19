package lcweekend

func longestBalanced(s string) int {
	if len(s) == 1 {
		return 0
	}

	ans := 0

	arr := make([]int, len(s))
	zero_cnt, one_cnt := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero_cnt++
			arr = append(arr, -1)
		} else {
			one_cnt++
			arr = append(arr, 1)
		}
	}
	if zero_cnt == 0 || one_cnt == 0 {
		return ans
	}

	return ans

}
