package linklist

// https://leetcode.cn/problems/odd-even-linked-list/
// 奇数 偶数 指针
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	f, s := head, head.Next
	ff, ss := f, s
	for s != nil && s.Next != nil {
		f.Next = s.Next
		s.Next = s.Next.Next
		f = f.Next
		s = s.Next
	}
	f.Next = ss
	return ff
}
