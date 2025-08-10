package tree

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
// 层次遍历，空值也入队，当出队值为空时判断队列剩下元素是否都为空，为则完全二叉树
// 958 二叉树的完全性检验
func isCompleteTree(root *TreeNode) bool {
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		nums := len(nodeList)
		for nums > 0 {
			node := nodeList[0]

			if node == nil {
				for i := 0; i < len(nodeList); i++ {
					if nodeList[i] != nil {
						return false
					}
				}
				return true
			}

			nodeList = append(nodeList, node.Left)
			nodeList = append(nodeList, node.Right)

			nodeList = nodeList[1:]
			nums--
		}
	}
	return true
}
