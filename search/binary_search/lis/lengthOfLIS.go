package binarysearch

import "sort"

// https://leetcode.cn/problems/longest-increasing-subsequence/
// 300. 最长递增子序列

// 1. 创建一个数组a, 从原数组中获取数值p
// 2. 如果当前数值p比a的最后一位大, 那就直接放到a末尾
// 3. 如果当前数值p不比最后一位大, 用p覆盖掉比它大的元素中最小的那个;
func lengthOfLIS(nums []int) int {
	g := []int{}
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { // >=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}

// https://leetcode.cn/problems/longest-increasing-subsequence/solutions/13118/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-e/
// 原因：当用较小值覆盖 arr[j] 时，这个较小值在原数组中可能出现在 arr 后续元素之后，破坏了子序列的顺序约束。这个算法本质上只维护了"长度为 j+1
// 的递增子序列的最小可能末尾值"，而不是子序列本身。
func lengthOfLIS_1(nums []int) int {
	arr := []int{nums[0]}

	n := len(nums)
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > arr[len(arr)-1] {
			arr = append(arr, num)
			continue
		}
		// 与普通二分查找不同的是, 这个题目并不是查找具体的数据,
		// 而是找到应该插入的位置, 所有大于num的最左边
		i, j := 0, len(arr)-1
		for i <= j {
			mid := i + (j-i)/2
			if arr[mid] >= num {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
		arr[i] = num
	}
	return len(arr)
}
