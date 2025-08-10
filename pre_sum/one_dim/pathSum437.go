package onedim

import "fmt"

// https://leetcode.cn/problems/path-sum-iii/
// 额外打印路径
func pathSum(root *TreeNode, targetSum int) int {

	// 数据初始化
	preSumCntMap := map[int]int{0: 1}
	ans := 0

	var dfs_pathSum func(*TreeNode, int, []int)
	dfs_pathSum = func(node *TreeNode, preSum int, curPath []int) {
		if node == nil {
			return
		}

		// 更新路径, 前缀和
		curPath = append(curPath, node.Val)
		curPreSum := preSum + node.Val

		// 记录结果
		ans += preSumCntMap[curPreSum-targetSum]
		// 检查所有以当前节点结尾的路径
		if preSumCntMap[curPreSum-targetSum] > 0 {
			sum := 0
			for i := len(curPath) - 1; i >= 0; i-- {
				sum += curPath[i]
				if sum == targetSum {
					fmt.Println(curPath[i:])
				}
			}
		}

		preSumCntMap[curPreSum]++ // 计算完之后,才计数加1
		dfs_pathSum(node.Left, curPreSum, curPath)
		dfs_pathSum(node.Right, curPreSum, curPath)

		// 回溯
		curPath = curPath[:len(curPath)-1]
		preSumCntMap[curPreSum]--
	}

	dfs_pathSum(root, 0, []int{})

	return ans
}
