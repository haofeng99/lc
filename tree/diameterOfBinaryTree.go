package tree

// https://leetcode.cn/problems/diameter-of-binary-tree/
// 543 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	dfs_diameterOfBinaryTree(root)
	return maxDiameter
}

var maxDiameter int

func dfs_diameterOfBinaryTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftDiameter := dfs_diameterOfBinaryTree(node.Left)
	rightDiameter := dfs_diameterOfBinaryTree(node.Right)

	temp := leftDiameter + rightDiameter + 1
	if temp > maxDiameter {
		maxDiameter = temp
	}
	return max(leftDiameter, rightDiameter) + 1
}

var maxDiameter_path []int

func dfs_diameterOfBinaryTree_way(node *TreeNode) (int, []int) {
	if node == nil {
		return 0, []int{}
	}

	leftDiameter, leftArr := dfs_diameterOfBinaryTree_way(node.Left)
	rightDiameter, rightArr := dfs_diameterOfBinaryTree_way(node.Right)

	temp := leftDiameter + rightDiameter + 1
	if temp > maxDiameter {
		maxDiameter = temp

		temp := []int{}
		temp = append(temp, leftArr...)
		temp = append(temp, node.Val)
		temp = append(temp, rightArr...)

		maxDiameter_path = temp
	}

	if leftDiameter >= rightDiameter {
		return leftDiameter + 1, append(leftArr, node.Val)
	} else {
		// right 反转
		return rightDiameter + 1, append([]int{node.Val}, rightArr...)
	}
}
