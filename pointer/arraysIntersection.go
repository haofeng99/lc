package pointer

// 1213. 三个有序数组的交集
// 给出三个均为严格递增排列的正数数组arr1, arr2和arr3. 返回一个由仅在这三个数组中同时出现的整数所构成数组

// 思路: 三指针
func arraysIntersection(arr1, arr2, arr3 []int) []int {
	ans := []int{}
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		if arr1[i] == arr2[j] && arr1[i] == arr3[k] {
			ans = append(ans, arr1[i])
			i++
			j++
			k++
		} else {
			// 找到三个数据中最小的
			min := min(min(arr1[i], arr2[j]), arr3[k])
			if min == arr1[i] {
				i++
			}
			if min == arr2[j] {
				j++
			}
			if min == arr3[k] {
				k++
			}
		}
	}

	return ans
}

func TestArraysIntersection(arr1, arr2, arr3 []int) []int {
	return arraysIntersection(arr1, arr2, arr3)
}
