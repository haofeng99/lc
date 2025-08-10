package linklist

// https://leetcode.cn/problems/remove-duplicate-node-lcci/
func removeDuplicateNodes(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		v := cur.Val
		prev := cur
		for prev != nil && prev.Next != nil {
			if prev.Next.Val == v {
				prev.Next = prev.Next.Next
			} else {
				prev = prev.Next
			}
		}
		cur = cur.Next
	}
	return head
}

func removeDuplicateNodesII(head *ListNode) *ListNode {
	dupMap := make(map[int]int)

	cur := head
	idx := 0
	for cur != nil {
		if _, ok := dupMap[cur.Val]; !ok {
			dupMap[cur.Val] = idx
		}
		cur = cur.Next
		idx++
	}

	dummy := &ListNode{-1, head}
	pre := dummy
	cur = head
	idx = 0
	for cur != nil {
		if dupMap[cur.Val] == idx {
			pre.Next = cur
			pre = pre.Next
		}
		cur = cur.Next
		idx++
	}
	pre.Next = nil
	return dummy.Next
}
