package linklist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{-101, head}
	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return dummy.Next
}
