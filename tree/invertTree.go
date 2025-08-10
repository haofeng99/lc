package tree

// https://leetcode.cn/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func dfs_invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		nums := len(nodeList)
		for nums > 0 {
			node := nodeList[0]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
			nodeList = nodeList[1:]
			nums--
		}
	}
	return root
}
