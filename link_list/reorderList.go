package linklist

// 1.先找链表中间点,并分割为前后两个子链表
// 2.反转第二个子链表
// 3.交叉合并两个子链表
func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondListHead := slow.Next
	slow.Next = nil
	head2 := reserveSecondList(secondListHead)
	crossMerge(head, head2)
}

func crossMerge(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{-1, h1}
	cur := dummy
	for h1 != nil || h2 != nil {
		if h1 != nil {
			cur.Next = h1
			h1 = h1.Next
			cur = cur.Next
		}
		if h2 != nil {
			cur.Next = h2
			h2 = h2.Next
			cur = cur.Next
		}
	}
	return dummy.Next
}

func reserveSecondList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}
