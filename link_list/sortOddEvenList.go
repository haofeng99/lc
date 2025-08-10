package linklist

// https://mp.weixin.qq.com/s/0WVa2wIAeG0nYnVndZiEXQ
// 给定一个奇数位升序，偶数位降序的链表，将其重新排序
// 1->8->3->6->5->4->7->2->NULL

// 思路:
func sortOddEventList(head *ListNode) *ListNode {
	h1, h2 := _splitList(head)
	newH2 := _reserveList(h2)
	return _mergeTwoSortedList(h1, newH2)
}

func _mergeTwoSortedList(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	dummy := &ListNode{-1, nil}
	cur := dummy
	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return dummy.Next
}

func _reserveList(h *ListNode) *ListNode {
	if h == nil {
		return nil
	}
	var pre *ListNode
	cur := h
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

func _splitList(h *ListNode) (*ListNode, *ListNode) {
	if h == nil {
		return nil, nil
	}
	h1, h2 := h, h.Next
	dummy1, dummy2 := &ListNode{0, h1}, &ListNode{0, h2}
	for h2 != nil && h2.Next != nil {
		cur := h2.Next
		h1.Next = cur
		h2.Next = cur.Next
		h1, h2 = cur, cur.Next
	}
	h1.Next = nil

	return dummy1.Next, dummy2.Next
}
