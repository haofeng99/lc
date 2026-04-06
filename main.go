package main

import (
	"fmt"
)

func main() {

	// 1. slice初始化
	// 1.1 make初始化
	arr_make := make([]int, 5, 10)
	fmt.Println(arr_make) // [0 0 0 0 0] make初始化会将元素初始化为零值

	arr_make_0 := make([]int, 0, 10)
	// fmt.Println(arr_make_0[0]) // panic: runtime error: index out of range [0] with length 0
	arr_make_0 = append(arr_make_0, 1)
	fmt.Println(arr_make_0) // [1]

	// 1.2 new初始化(通常不建议使用)
	// new函数用于创建指定类型的零值对象，并返回该对象的指针
	arr_new := new([]int)
	fmt.Println(arr_new) // &[] 这是Go的fmt包对*[]T类型的特殊打印格式
	// arr_new = append(arr_new, 1) // first argument to append must be a slice; have arr_new (variable of type *[]int)

	arr_new_5 := new([5]int)
	fmt.Println(arr_new_5) //&[0 0 0 0 0]

	// 1.3 直接声明
	arr_literal := []int{}
	fmt.Println(arr_literal) // []
	// fmt.Println(arr_literal[0]) // panic: runtime error: index out of range [0] with length 0

	// 1.4 字面量
	var arr []int
	fmt.Println(arr) // [] 字面量初始化会将元素初始化为零值
	// fmt.Println(arr[0]) // panic: runtime error: index out of range [0] with length 0

	// 1.5 二维slice初始化
	matrix := make([][]int, 3)
	fmt.Println(matrix) // [[] [] []]
	// matrix[0][0] = 1 // panic: runtime error: index out of range [0] with length 0
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 1
	fmt.Println(matrix) // [[1 0 0] [0 0 0] [0 0 0]]

	// 1.6 二维矩阵初始化1
	matrix_2 := make([][2]int, 3)
	fmt.Println(matrix_2) // [[0 0] [0 0] [0 0]]

	// 2. map初始化
}
