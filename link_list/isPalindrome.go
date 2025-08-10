package linklist

// https://leetcode.cn/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	h2 := slow.Next
	slow.Next = nil

	h2 = reverseList(h2)

	for head != nil && h2 != nil {
		if head.Val != h2.Val {
			return false
		}
		head = head.Next
		h2 = h2.Next
	}

	return true
}
