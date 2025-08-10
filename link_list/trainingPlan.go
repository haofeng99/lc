package linklist

// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := head, head
	for cnt > 0 {
		fast = fast.Next
		cnt--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
