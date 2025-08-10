package linklist

// https://leetcode.cn/problems/linked-list-cycle-ii/
// 1.快慢指针 确定有没有环路
// 2.有环路时, 快指针从新回到头节点, 并且快指针和慢指针同速度, 然后下一次快慢指针相遇点即为入环口
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
