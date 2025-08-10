package twodim

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {

	h, w := len(matrix), len(matrix[0])
	preSum := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		preSum[i] = make([]int, w+1)
	}

	// 计算二维前缀和
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + int(matrix[i-1][j-1]) - 48
		}
	}

	res := 0
	l := 1
	for l <= min(h, w) {
		for i := l; i <= h; i++ {
			for j := l; j <= w; j++ {
				if preSum[i][j]-preSum[i-l][j]-preSum[i][j-l]+preSum[i-l][j-l] == l*l {
					res = max(res, l)
				}
			}
		}
		l++
	}

	return res * res
}

func Show_maxSquare() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	ans := maximalSquare(matrix)
	fmt.Println(ans)
}
