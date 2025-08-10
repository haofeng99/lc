package sort

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	// k桶下标 v桶内元素
	bucketMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx := getIdx(nums[i], valueDiff+1)
		if _, ok := bucketMap[idx]; ok {
			return true
		}
		if v, ok := bucketMap[idx+1]; ok {
			if abs(v-nums[i]) <= valueDiff {
				return true
			}
		}
		if v, ok := bucketMap[idx-1]; ok {
			if abs(v-nums[i]) <= valueDiff {
				return true
			}
		}
		bucketMap[idx] = nums[i]
		if i >= indexDiff {
			delete(bucketMap, getIdx(nums[i-indexDiff], valueDiff+1))
		}
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

func getIdx(num, size int) int {
	if num >= 0 {
		return num / size
	}
	return (num+1)/size - 1
}

// https://leetcode.cn/problems/contains-duplicate-iii/solutions/726905/gong-shui-san-xie-yi-ti-shuang-jie-hua-d-dlnv/

// 1.为什么桶的大小需要是valueDiff+1,  因为，所以需要+1
// 令size = valueDiff + 1 的本质是因为差值为 valueDiff 两个数在数轴上相隔距离为 valueDiff + 1，它们需要被落到同一个桶中。
// 如果桶的大小是valueDiff的话，那么如果存在两个数差值是valueDiff，那么这两个数肯定不在同一个桶里
// 例如valueDiff = 1 且 桶size设置为1的话，那么差值为1的两个数，比如1,2, 在数轴上相隔距离为2，且为2个数, 就没办法落到同一个桶中, 而根据解题语义它们需要被落到同一个桶中。

// getIdx需要分别考虑正数和负的情况
// 1.正数时直接使用数据除以桶大小,就可以得到数会落到哪个桶;
// 2.负数时 https://leetcode.cn/problems/contains-duplicate-iii/solutions/726905/gong-shui-san-xie-yi-ti-shuang-jie-hua-d-dlnv/

// 为什么不用担心一个桶内元素重复的情况
// 因为如果桶已经存在,则说明满足情况,直接return true

// 桶排序的本质就是将一个数映射到一个具体的位置[bucket]
