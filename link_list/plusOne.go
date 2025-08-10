package linklist

// https://www.nowcoder.com/questionTerminal/a2f1105d2ac5466e9ba8fd61310ba6d1
func plusOne(head *ListNode) *ListNode {
	// write code here
	flag := 0
	head = _reserveList(head)
	cur := head
	sum := 1
	for cur != nil {
		sum += cur.Val + flag
		flag = sum / 10
		cur.Val = sum % 10
		sum = 0
		cur = cur.Next
	}
	head = _reserveList(head)
	if flag != 0 {
		node := &ListNode{flag, head}
		head = node
	}
	return head
}
