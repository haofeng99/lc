package binarysearchnosort

import (
	"fmt"
)

const Mod = int(1e9 + 7)

func maxProfit(inventory []int, orders int) int {
	max := inventory[0]
	for _, in := range inventory {
		if in > max {
			max = in
		}
	}

	l, r := 0, max-1
	k := -1
	for l <= r {
		mid := l + (r-l)>>1
		if count(inventory, mid) > orders {
			l = mid + 1
		} else if count(inventory, mid) < orders {
			r = mid - 1
		} else {
			k = mid
			break
		}
	}
	if k == -1 {
		k = l
	}
	ans := 0
	for _, in := range inventory {
		if in > k {
			orders = orders - (in - k)
			ans = (ans + sum(in, k)) % Mod
		}
	}
	ans = (ans + k*orders) % Mod
	return ans
}

func sum(top, lower int) int {
	return ((top + lower + 1) * (top - lower) / 2) % Mod
}

// 统计可以卖的个数
// x 大于x的可以卖
func count(inventory []int, x int) int {
	sum := 0
	for _, i := range inventory {
		if i > x {
			sum = (sum + i - x) % Mod
		}
	}
	return sum
}

func Show_maxProfit() {
	inventory := []int{
		497978859,
		167261111,
		483575207,
		591815159}
	orders := 836556809
	ans := maxProfit(inventory, orders)
	fmt.Println(ans)
}
