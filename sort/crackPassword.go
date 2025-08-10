package sort

import (
	"fmt"
	"sort"
	"strconv"
)

func crackPassword(password []int) string {
	sort.Slice(password, func(i, j int) bool {
		px, py := 10, 10
		x, y := password[i], password[j]
		for px <= y {
			px *= 10
		}
		for py <= x {
			py *= 10
		}
		return x*px+y < y*py+x
	})
	ans := []byte{}
	for _, x := range password {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func Show_crackPassword() {
	password := []int{2, 10}
	fmt.Println(crackPassword(password))
}
