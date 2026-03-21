package hash

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// LCR 120. 寻找文件副本

func findRepeatDocument(documents []int) int {
	for i := 0; i < len(documents); i++ {
		documents[i]++
	}
	for i, d := range documents {
		if d < 0 {
			d = -d
		}
		if documents[d-1] < 0 {
			return i
		}
		documents[d-1] = -documents[d-1]
	}
	return -1
}
