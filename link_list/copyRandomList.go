package linklist

// https://leetcode.cn/problems/copy-list-with-random-pointer/
// A  ->  B -> C
// A  ->  A1 -> B -> B1 -> C -> C1
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	// 复制
	cur := head
	for cur != nil {
		curCopy := &Node{cur.Val, cur.Next, cur.Random}
		cur.Next = curCopy
		cur = curCopy.Next
	}

	// Random重新指向
	cur = head
	for cur != nil {
		curRandom := cur.Random
		if curRandom != nil {
			cur.Next.Random = curRandom.Next
		}
		cur = cur.Next.Next
	}

	// 链表拆分, 原始链表复原
	cur = head
	newHead := head.Next
	for cur != nil {
		curCopy := cur.Next
		curNext := curCopy.Next

		cur.Next = curNext

		if curNext != nil {
			curCopy.Next = curNext.Next
		}
		cur = curNext
	}

	return newHead
}

// 两次遍历
// 不通过,当同时修改next和random节点时,random可能会指向旧链表,可以通过额外开辟空间map解决
func copyRandomListII(head *Node) *Node {
	if head == nil {
		return head
	}

	// 复制
	cur := head
	for cur != nil {
		curCopy := &Node{cur.Val, cur.Next, cur.Random}
		cur.Next = curCopy
		cur = curCopy.Next
	}

	// 重新指向
	cur = head
	newHead := head.Next
	for cur != nil {

		curCopy := cur.Next

		// Random指针
		curRandom := cur.Random
		if curRandom != nil {
			curCopy.Random = curRandom.Next
		}

		// Next指针
		curCopyNext := curCopy.Next
		if curCopyNext != nil {
			curCopy.Next = curCopyNext.Next
		}

		// 原链表复原
		cur.Next = curCopyNext

		cur = cur.Next
	}

	return newHead
}
