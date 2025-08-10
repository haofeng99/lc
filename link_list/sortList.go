package linklist

// 归并排序写法

// https://leetcode.cn/problems/sort-list/description/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return splitAndMerge(head)
}

func splitAndMerge(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	midNode := getMid(head)
	rightListHead := midNode.Next

	midNode.Next = nil
	leftList := head

	return mergeTwoSortedList(splitAndMerge(leftList), splitAndMerge(rightListHead))
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoSortedList(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return dummy.Next
}
