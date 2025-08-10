package tree

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
// 107. 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nums := len(queue)
		levelArr := []int{}
		for nums > 0 {
			node := queue[0]
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
			nums--
		}
		ans = append([][]int{levelArr}, ans...)
	}
	return ans
}
