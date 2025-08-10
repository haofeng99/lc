package tree

// 450-删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/
// 1.寻找待删除节点
// 1.1 如果当前节点值大于目标节点, 说明待删除节点在左子树, 向左子树递归
// 1.2 如果当前节点值小于目标节点, 说明待删除节点在右子树, 向右子树递归
// 1.3 如果当前节点值等于目标节点, 说明当前节点就是要删除节点, 走删除逻辑
// 2.删除逻辑
// 2.1 如果当前节点是子节点, 则直接return nil
// 2.2 若干当前节点左右子节点只同时存在一个, 则直接用存在的左或右子节点替代当前节点
// 3.3 若干当前节点左右子节点都存在, 为了删除当前节点, 需要找到大于当前节点的最小节点(右子树的最左节点)替换当前节点
//
//	然后向右子树递归, 删除最小节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// 找到大于当前值的最小值
			// 右子树的最左节点
			next := root.Right
			for next.Left != nil {
				next = next.Left
			}
			root.Val = next.Val
			root.Right = deleteNode(root.Right, root.Val) // 为什么是root.Val而不是key? 因为这里是要迭代删除next
			return root
		}
	}
}

// 迭代思路:
// 递归主要是在找需要删除的目标节点, 这部分可以通过迭代实现
func deleteNode_bfs(root *TreeNode, key int) *TreeNode {
	var parent, cur *TreeNode = nil, root
	for cur != nil && cur.Val != key {
		parent = cur
		if cur.Val < key {
			cur = cur.Right
		} else if cur.Val > key {
			cur = cur.Left
		}
	}
	// 没找到待删除的节点
	if cur == nil {
		return root
	}

	del := func(node *TreeNode) *TreeNode {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			next, nextParent := node.Right, node
			for next.Left != nil {
				nextParent = next
				next = next.Left
			}
			node.Val = next.Val // 删除[覆盖]node节点
			// 删除next节点 (上面的循环可能不执行, 所以nextParnet和next的关系不确定)
			if nextParent.Left == next {
				nextParent.Left = next.Right
			} else {
				nextParent.Right = next.Right
			}
			return node
		}
	}

	// 要删除的是根节点
	if parent == nil {
		return del(cur)
	}
	// 要删除的是左子节点
	if parent.Left == cur {
		parent.Left = del(cur)
	} else {
		// 要删除右子节点
		parent.Right = del(cur)
	}
	return root
}
