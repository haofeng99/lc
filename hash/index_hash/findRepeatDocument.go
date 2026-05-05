package hash

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// LCR 120. 寻找文件副本

// 方法1: 为了避免0数据的出现, 直接将原有数据加n, 重新映射到n -> 2n之间
func findRepeatDocument(documents []int) int {
	n := len(documents)
	for _, v := range documents {
		num := v % n // 重新映射回0～n-1
		if documents[num] >= n {
			return num
		}
		documents[num] += n
	}
	return -1
}

// 方法2: 交换法
// 把每个数交换到它"应该在"的索引位置，如果目标位置已经有正确的数，则找到重复
// 不会出现死循环
func findRepeatDocument_ii(documents []int) int {
	n := len(documents)
	for i := 0; i < n; i++ {
		for documents[documents[i]] != documents[i] {
			documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
		}
		if documents[i] != i {
			return documents[i]
		}
	}
	return -1
}
