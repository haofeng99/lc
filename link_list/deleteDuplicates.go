package linklist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
// pre, cur, next 三节点
// cur 和 next节点值不等, 三节点正常向后移动
// cur和next节点值相等, 一直向后找next节点,直到和cur节点值不等, 然后 pre.Next = next;  cur = next 分情况分析next是否为空
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{-101, head}
	pre, cur, next := dummy, head, head.Next
	for next != nil {
		hasDup := false
		for next != nil && cur.Val == next.Val {
			hasDup = true
			next = next.Next
		}
		if !hasDup {
			pre, cur, next = cur, next, next.Next
		} else {
			pre.Next = next
			cur = next
			if next != nil {
				next = next.Next
			}
		}

	}
	return dummy.Next
}
