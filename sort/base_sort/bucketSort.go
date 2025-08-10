package basesort

import "math"

func BucketSort(arr []int) []int {
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	bucketNum := int(math.Sqrt(float64(len(arr))))
	bucket := make([][]int, bucketNum)
	for j := 0; j < len(arr); j++ {
		// 分桶公式
		n := bucketNum * (arr[j] - min) / (max - min + 1)
		bucket[n] = append(bucket[n], arr[j])
		// 桶内单轮次排序
		k := len(bucket[n]) - 2
		for k >= 0 && bucket[n][k+1] < bucket[n][k] {
			bucket[n][k], bucket[n][k+1] = bucket[n][k+1], bucket[n][k]
			k--
		}
	}

	i := 0
	for _, v := range bucket {
		for _, vv := range v {
			arr[i] = vv
			i++
		}
	}
	return arr
}
