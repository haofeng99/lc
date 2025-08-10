package linklist

// https://leetcode.cn/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	sH := &ListNode{0, head}
	bH := &ListNode{0, head}
	s, b := sH, bH
	cur := head
	for cur != nil {
		if cur.Val < x {
			s.Next = cur
			s = s.Next
		} else {
			b.Next = cur
			b = b.Next
		}
		cur = cur.Next
	}
	b.Next = nil
	s.Next = bH.Next
	return sH.Next
}
