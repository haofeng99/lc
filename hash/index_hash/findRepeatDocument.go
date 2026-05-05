package hash

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// LCR 120. 寻找文件副本

// 方法1: 0的出现会导致无法取负数,
// 因此不取负数, 而是往上取大数,
// 在每次遍历时, 取大数的模, 还原来的数, 并判断对应下标位置的数据是否为大数
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
		// 推出内层for循环后, 状态为documents[documents[i]] != documents[i]
		// 那么此时, 如果documents[i] != i, 就代表两个不同的位置出现了相同的数据
		if documents[i] != i {
			return documents[i]
		}
	}
	return -1
}
