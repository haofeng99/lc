package linklist

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	idx := 0
	cache := make(map[int]*ListNode)
	for ptr != nil {
		cache[idx] = ptr
		ptr = ptr.Next
		idx++
	}
	p := idx - n
	pre := p - 1
	if pre < 0 {
		return cache[1]
	}
	next := p + 1
	if next >= idx {
		cache[p-1].Next = nil
		return cache[0]
	}
	cache[p-1].Next = cache[p+1]
	return head
}

// 栈
// 快慢指针
