package tree

// https://leetcode.cn/problems/binary-tree-level-order-traversal/

// bfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	deep := 0
	nodeList := []*TreeNode{root}
	for len(nodeList) != 0 {
		ans = append(ans, []int{})
		nodeSize := len(nodeList)
		for nodeSize > 0 {
			ans[deep] = append(ans[deep], nodeList[0].Val)
			node := nodeList[0]
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
			nodeSize--
			nodeList = nodeList[1:]
		}
		deep++
	}

	return ans
}

// dfs
func levelOrder_(root *TreeNode) [][]int {
	ans := [][]int{}
	dfs(root, &ans, 0)
	return ans
}

func dfs(node *TreeNode, ans *[][]int, deep int) {
	if node == nil {
		return
	}
	if deep+1 > len(*ans) {
		*ans = append(*ans, []int{})
	}
	if deep%2 == 0 {
		// 从左向右放
		(*ans)[deep] = append((*ans)[deep], node.Val)
	} else {
		// 从右向左放
		nodeVal := []int{node.Val}
		nodeVal = append(nodeVal, (*ans)[deep]...)
		(*ans)[deep] = nodeVal
	}
	dfs_zigzagLevelOrder(node.Left, ans, deep+1)
	dfs_zigzagLevelOrder(node.Right, ans, deep+1)
}
