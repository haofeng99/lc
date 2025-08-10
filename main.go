package main

import (
	"fmt"
	"math"
)

func main() {

	str := "[aaa]"
	fmt.Println(str[1 : len(str)-1])

	// cache5 := map[int][][]int{}
	// cache5[1][0] = append(cache5[1][0], 55)
	// fmt.Println(cache5)

	// cache4 := [][]int{}
	// cache4[0] = append(cache4[0], 5)
	// fmt.Println(cache4)

	// cache := make(map[int][]int)
	// cache[1] = append(cache[1], 2)
	// cache[1] = append(cache[1], 3)
	// fmt.Println(cache)

	// cache2 := []int{}
	// cache2 = append(cache2, 4)
	// fmt.Println(cache2)

	// ans1 := [][]int{}
	// ans1[0] = append(ans1[0], 5)
	// fmt.Println(ans1)

	// stack := make([]int, 5)
	// fmt.Println(stack)

	// cache1 := map[int][]int{}
	// cache1[0] = append(cache1[0], 1)
	// fmt.Println(cache1)

	// cache1 := make(map[int][][]int)
	// cache1[1] = append(cache1[1], []int{1, 2, 3})
	// cache1[1] = append(cache1[1], []int{3, 4, 5})
	// fmt.Println(cache1)

	// cache2 := map[int][][]int{}
	// cache2[1] = append(cache2[1], []int{1, 2, 3})
	// cache2[1] = append(cache2[1], []int{3, 4, 5})
	// fmt.Println(cache2)

	// ans2 := make([][]int, 5)
	// ans2[1] = append(ans2[1], 5)
	// fmt.Println(ans2)
	fmt.Println(math.Pow(2, 3))

	s := "->aaa"
	fmt.Println(s[0:2])
	fmt.Println(s[2:])
}
