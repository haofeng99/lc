package divideconquer

import (
	"strconv"
)

// 241. 为运算表达式设计优先级
// 思路: https://leetcode.cn/problems/different-ways-to-add-parentheses/solutions/44108/pythongolang-fen-zhi-suan-fa-by-jalan/
// 1.按运算符分割左右两部分, 分别求解
// 2.针对左边结果和右边结果, 结合当前操作符, 得到所有结果
func diffWaysToCompute(expression string) []int {
	var isDigit func(string) bool
	isDigit = func(str string) bool {
		_, err := strconv.Atoi(str)
		return err == nil
	}

	var calc func([]int, []int, string) []int
	calc = func(v1, v2 []int, op string) []int {
		ans := []int{}
		for _, a := range v1 {
			for _, b := range v2 {
				if op == "+" {
					ans = append(ans, a+b)
					continue
				}
				if op == "-" {
					ans = append(ans, a-b)
					continue
				}
				ans = append(ans, a*b)
			}
		}
		return ans
	}

	if isDigit(expression) {
		v, _ := strconv.Atoi(expression)
		return []int{v}
	}

	ans := []int{}
	for idx, r := range expression {
		char := string(r)
		if char == "+" || char == "-" || char == "*" {
			// 左边可能的值与右边可能的值
			left := diffWaysToCompute(expression[:idx])
			right := diffWaysToCompute(expression[idx+1:])

			calc_ans := calc(left, right, char)
			ans = append(ans, calc_ans...)
		}
	}

	return ans
}
