package linklist

// 方法1: 找到left左边的节点pre和right右边的节点next，然后反转left到right之间的节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	cur := &ListNode{-1, head}
	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}

	leftNode := cur
	endNode := cur.Next
	cur = cur.Next
	for i := left; i < right; i++ {
		cur = cur.Next
	}
	rightNode := cur.Next
	startNode := cur

	startNode.Next = nil
	leftNode.Next = nil

	e := reverseListHelper(endNode)

	leftNode.Next = e
	endNode.Next = rightNode

	return head
}

// 反转链表
func reverseListHelper(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre, curr = curr, next
	}
	return pre
}

// 方法2: 栈
func reverseBetweenII(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{-1, head}
	pos := 0
	cur := dummy
	for pos < left-1 {
		cur = cur.Next
		pos++
	}
	pre := cur

	pos++
	cur = cur.Next

	stack := make([]*ListNode, 0)
	for pos <= right {
		stack = append(stack, cur)
		cur = cur.Next
		pos++
	}

	next := cur

	for len(stack) > 0 {
		pre.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre = pre.Next
	}
	pre.Next = next

	return dummy.Next
}

// 如果 left 和 right 的区域很大，恰好是链表的头节点和尾节点时，找到 left 和 right 需要遍历一次，
// 反转它们之间的链表还需要遍历一次，虽然总的时间复杂度为 O(N)，但遍历了链表 2 次，

// 一次遍历-穿针引线法
func reverseBetweenIII(head *ListNode, left int, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{-1, head}
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	// n个节点 只需要移动n-1次
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		// 为什么这里不是cur, 因为需要将next节点放入pre节点的后面,而不是cur节点的前面
		// 因为cur是会一直向后移动的,而每次next接节点都是直接放到pre节点的后面的固定位置
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

// https://leetcode.cn/problems/reverse-linked-list-ii/
