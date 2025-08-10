package linklist

// https://leetcode.cn/problems/merge-in-between-linked-lists/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, list1}
	pre, next := find(list1, a, b)
	pre.Next = list2

	cur := list2
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = next

	return dummy.Next
}

func find(list1 *ListNode, a int, b int) (*ListNode, *ListNode) {
	dummy := &ListNode{-1, list1}
	idx := 0
	pre, cur := dummy, list1
	var ans1, ans2 *ListNode
	for cur != nil {
		if idx == a {
			ans1 = pre
		}
		if idx == b {
			ans2 = cur.Next
		}
		if ans1 != nil && ans2 != nil {
			return ans1, ans2
		}
		pre = pre.Next
		cur = cur.Next
		idx++
	}
	return ans1, ans2
}
