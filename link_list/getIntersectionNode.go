package linklist

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cache := make(map[*ListNode]bool)
	for headA != nil {
		cache[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if cache[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
