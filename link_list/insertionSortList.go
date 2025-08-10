package linklist

// https://leetcode.cn/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{5001, head}
	cur := dummy
	for cur != nil {
		minValPre := findMinValPreNode(cur)
		if minValPre != nil && minValPre != cur {
			cur.Next, minValPre.Next = minValPre.Next, cur.Next
			cur.Next.Next, minValPre.Next.Next = minValPre.Next.Next, cur.Next.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 从某个位置开始, 找到最小元素的前一个元素
func findMinValPreNode(startNode *ListNode) *ListNode {
	if startNode.Next == nil {
		return nil
	}
	minVal := startNode.Next.Val
	cur := startNode.Next
	pre := startNode
	ans := pre
	for cur != nil {
		if cur.Val < minVal {
			minVal = cur.Val
			ans = pre
		}
		cur = cur.Next
		pre = pre.Next
	}
	return ans
}
