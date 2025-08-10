package tree

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	nodeList := []*TreeNode{root}
	direction := 1
	for len(nodeList) > 0 {
		nodeSize := len(nodeList)
		curArr := []int{}
		for i := 0; i < nodeSize; i++ {
			node := nodeList[i]
			if direction == 1 {
				curArr = append(curArr, node.Val)
			} else {
				temp := []int{node.Val}
				temp = append(temp, curArr...)
				curArr = temp
			}
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
		}
		ans = append(ans, curArr)
		direction = direction * -1
		nodeList = nodeList[nodeSize:]
	}
	return ans
}

func zigzagLevelOrder_(root *TreeNode) [][]int {
	ans := [][]int{}
	dfs_zigzagLevelOrder(root, &ans, 0)
	return ans
}

func dfs_zigzagLevelOrder(node *TreeNode, ans *[][]int, deep int) {
	if node == nil {
		return
	}
	if deep+1 > len(*ans) {
		*ans = append(*ans, []int{})
	}
	curDeepArr := (*ans)[deep]
	if deep == 0 || deep%2 == 0 {
		// 从左向右放
		curDeepArr = append(curDeepArr, node.Val)
	} else {
		// 从右向左放
		nodeVal := []int{node.Val}
		nodeVal = append(nodeVal, curDeepArr...)
		(*ans)[deep] = nodeVal
	}
	dfs_zigzagLevelOrder(node.Left, ans, deep+1)
	dfs_zigzagLevelOrder(node.Right, ans, deep+1)
}
