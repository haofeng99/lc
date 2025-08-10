package tree

// https://leetcode.cn/problems/merge-two-binary-trees/
// 617. 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue1 := []*TreeNode{root1}
	queue2 := []*TreeNode{root2}
	curParent := &TreeNode{Val: -1}
	queue3 := []*TreeNode{curParent}

	idx := 1
	for len(queue1) != 0 && len(queue2) != 0 {
		node1, node2 := queue1[0], queue2[0]
		if node1 != nil || node2 != nil {
			if node1 == nil {
				node1 = &TreeNode{Val: 0}
			}
			if node2 == nil {
				node2 = &TreeNode{Val: 0}
			}
			node1.Val = node1.Val + node2.Val
			if idx%2 == 0 {
				queue3[0].Left = node1
			} else {
				queue3[0].Right = node1
			}
			queue3 = append(queue3, node1)
		}
		idx++
		if node1 != nil {
			queue1 = append(queue1, node1.Left)
			queue1 = append(queue1, node1.Right)
		}
		if node2 != nil {
			queue2 = append(queue2, node2.Left)
			queue2 = append(queue2, node2.Right)
		}
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if idx%2 == 0 {
			queue3 = queue3[1:]
		}
	}

	return curParent.Right
}

func mergeTrees_dfs(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees_dfs(root1.Left, root2.Left)
	root1.Right = mergeTrees_dfs(root1.Right, root2.Right)
	return root1
}
