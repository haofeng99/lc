package linklist

// https://leetcode.cn/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l := getLen(head)
	if k >= l {
		k = k % l
	}

	if k == 0 {
		return head
	}

	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	dummy := &ListNode{-1, slow}
	cur := slow
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = head
	cur = cur.Next

	for cur.Next != slow {
		cur = cur.Next
	}

	cur.Next = nil

	return dummy.Next

}
