package linklist

// https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{-1, head}
	pre, cur, next := dummy, head, head.Next
	for next != nil {
		nnext := next.Next
		pre.Next = next
		next.Next = cur
		cur.Next = nnext

		pre = cur
		cur = nnext
		if nnext != nil {
			next = nnext.Next
		} else {
			next = nil
		}
	}

	return dummy.Next
}

func swapPairsII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairsII(newHead.Next)
	newHead.Next = head
	return newHead
}
