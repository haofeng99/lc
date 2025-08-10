package linklist

// https://www.geeksforgeeks.org/problems/reverse-a-doubly-linked-list/1

func reverseDLL(head *DLLNode) *DLLNode {
	var pre *DLLNode
	for head != nil {
		next := head.Next

		head.Prev = next
		head.Next = pre

		pre = head
		head = next
	}

	return pre

}
