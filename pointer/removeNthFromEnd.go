package pointer

// 1.构建两个指针, 首先快指针先向后移动n次
// 2.此时两个指针同时同速向后移动,当快指针走到结尾时,慢指针所在的node就是要移除的node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	dummy := &ListNode{-1, head}
	pre := dummy
	for fast != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next
	}

	pre.Next = slow.Next
	return dummy.Next
}
