package sort

func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for _, num := range nums1 {
		map1[num]++
	}
	for _, num := range nums2 {
		map2[num]++
	}

	ans := make([]int, 0)
	for num, count := range map1 {
		if count2, ok := map2[num]; ok && count2 > 0 {
			for i := 0; i < min(count, count2); i++ {
				ans = append(ans, num)
			}
		}
	}

	return ans
}

// https://leetcode.cn/problems/intersection-of-two-arrays-ii/description/

// 如果给定的数组已经排好序呢？你将如何优化你的算法？
func intersect2(nums1 []int, nums2 []int) []int {
	p1, p2 := 0, 0
	ans := make([]int, 0)
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			num := nums1[p1]
			ans = append(ans, num)
			p1++
			p2++
		}
	}
	return ans
}

// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中,如何解决
func intersect3(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)

	for _, num := range nums1 {
		map1[num]++
	}

	ans := make([]int, 0)
	for _, num := range nums2 {
		// 每次匹配到一个,就减一个
		if count, ok := map1[num]; ok && count > 0 {
			ans = append(ans, num)
			map1[num]--
		}
	}

	return ans
}
