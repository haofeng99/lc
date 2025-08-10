package linklist

// https://leetcode.cn/problems/reverse-nodes-in-k-group/

// 思路: 反转start-end中间的k个链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}
	start := dummy
	end := start
	for end != nil {
		count := 0
		for count < k {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
			count++
		}
		start = reserveSubKList(start, end)
		end = start
	}

	return dummy.Next
}

// 反转(start-end]之间的子链表
// 返回新链表结尾
func reserveSubKList(start, end *ListNode) *ListNode {
	// start不参与排序 剥离
	head := start.Next
	start.Next = nil

	// end后续不参与排序 剥离
	endNext := end.Next
	end.Next = nil

	// 逆序head单链表
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}

	start.Next = pre
	head.Next = endNext
	return head
}
