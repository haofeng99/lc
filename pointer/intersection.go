package pointer

import "sort"

func intersection(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)

	p1, p2 := 0, 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] > nums2[p2] {
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			num := nums1[p1]
			if len(res) == 0 || res[len(res)-1] != num {
				res = append(res, num)
			}
			p1++
			p2++
		}
	}
	return res
}

func intersection2(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	res := make([]int, 0)
	for _, num := range nums1 {
		cache[num] = 1
	}
	for _, num := range nums2 {
		if count, ok := cache[num]; ok && count > 0 {
			res = append(res, num)
			delete(cache, num)
		}
	}
	return res
}
