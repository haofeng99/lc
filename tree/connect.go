package tree

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
// 117. 填充每个节点的下一个右侧节点指针 II

// O(n) O(n)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	var pre *Node
	for len(queue) > 0 {
		nums := len(queue)
		pre = nil
		for nums > 0 {
			node := queue[0]
			if pre != nil {
				pre.Next = node
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			pre = node
			queue = queue[1:]
			nums--
		}
	}

	return root
}

// O(n) O(1) 优化
// 渐进式链表构建
// 1.在第i层可直接完成第i+1层的链表构建
func connect_O1(root *Node) *Node {
	if root == nil {
		return root
	}

	cur := root
	for cur != nil {
		var nextLevelStart, curLevelPre *Node
		handle := func(node *Node) {
			if node == nil {
				return
			}
			if curLevelPre == nil {
				curLevelPre = node
				nextLevelStart = node
				return
			}
			curLevelPre.Next = node
			curLevelPre = node
		}

		for cur != nil {
			handle(cur.Left)
			handle(cur.Right)
			cur = cur.Next
		}

		cur = nextLevelStart
	}

	return root
}
