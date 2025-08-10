package linklist

import "fmt"

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre, curr = curr, next
	}
	return pre
}

func reverseListII(head *ListNode) *ListNode {
	var pre *ListNode
	return reverseListRecure(pre, head)
}

func reverseListRecure(pre, curr *ListNode) *ListNode {
	if curr == nil {
		return pre
	}
	next := curr.Next
	curr.Next = pre
	return reverseListRecure(curr, next)
}

// https://leetcode.cn/problems/reverse-linked-list/description/
func Show_ReverseList() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println("Original list:")
	PrintListNode(head)

	reversedHead := reverseList(head)
	fmt.Println("Reversed list:")
	PrintListNode(reversedHead)
}
