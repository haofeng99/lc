package sort

// https://leetcode.cn/problems/merge-k-sorted-lists/description
// 23. 合并 K 个升序链表
// 方法1: 堆
// 方法2: 归并排序
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeList(lists)
}

func mergeList(list []*ListNode) *ListNode {
	n := len(list)
	if n < 2 {
		return list[0]
	}
	mid := n / 2
	left := list[:mid]
	right := list[mid:]
	return mergeTwoList(mergeList(left), mergeList(right))
}

func mergeTwoList(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l != nil && r != nil {
		if l.Val <= r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}
	for l != nil {
		cur.Next = l
		l = l.Next
		cur = cur.Next
	}
	for r != nil {
		cur.Next = r
		r = r.Next
		cur = cur.Next
	}
	return dummy.Next
}
