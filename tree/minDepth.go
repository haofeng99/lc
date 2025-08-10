package tree

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 111-二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 1
	for len(queue) > 0 {
		levelLen := len(queue)
		for levelLen > 0 {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return ans
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			levelLen--
		}
		ans++
	}
	return ans
}
