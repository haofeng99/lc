package tree

// https://leetcode.cn/problems/symmetric-tree/  101-对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodeList := []*TreeNode{root}
	for len(nodeList) != 0 {
		num := len(nodeList)
		if nodeList[0] != root && num%2 != 0 {
			return false
		}
		curDeep := []int{}
		for num > 0 {
			node := nodeList[0]
			curDeep = append(curDeep, node.Val)
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			} else {
				if node.Val != -101 {
					nodeList = append(nodeList, &TreeNode{-101, nil, nil})
				}
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			} else {
				if node.Val != -101 {
					nodeList = append(nodeList, &TreeNode{-101, nil, nil})
				}
			}
			num--

			nodeList = nodeList[1:]
		}

		l, r := 0, len(curDeep)-1
		for l < r {
			if curDeep[l] != curDeep[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isSymmetric_dfs(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	// p 指针和 q 指针一开始都指向这棵树的根，随后 p 右移时，q 左移，p 左移时，q 右移。每次检查当前 p 和 q 节点的值是否相等，如果相等再判断左右子树是否对称
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root.Left, root.Right)
}

// https://leetcode.cn/problems/same-tree/  100-相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
func isSameTree_bfs(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		n1, n2 := queue[0], queue[1]
		if n1.Val != n2.Val {
			return false
		}
		left1, right1 := n1.Left, n1.Right
		left2, right2 := n2.Left, n2.Right
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}

		if left1 != nil {
			queue = append(queue, left1)
			queue = append(queue, left2)
		}
		if right1 != nil {
			queue = append(queue, right1)
			queue = append(queue, right2)
		}

		queue = queue[2:]
	}
	return true
}

// https://leetcode.cn/problems/subtree-of-another-tree/  572-另一棵树的子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/  LCR 143-子结构判断  字节-基础架构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		// p不可以为空, q可以为空
		if p == nil {
			return false
		}
		if q == nil {
			return true
		}
		return p.Val == q.Val && dfs(p.Left, q.Left) && dfs(p.Right, q.Right)
	}
	if dfs(A, B) {
		return true
	}

	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
