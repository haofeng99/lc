package linklist

// https://leetcode.cn/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	if getLen(l1) < getLen(l2) {
		l1, l2 = l2, l1
	}
	dummy := &ListNode{0, l1}
	cur := l1
	for l1 != nil {
		l2V := 0
		if l2 != nil {
			l2V = l2.Val
			l2 = l2.Next
		}
		l1V := l1.Val
		sum := l2V + l1V + flag
		if sum >= 10 {
			l1.Val = sum - 10
			flag = 1
		} else {
			l1.Val = sum
			flag = 0
		}
		cur = l1
		l1 = l1.Next
	}
	if flag == 1 {
		cur.Next = &ListNode{1, nil}
	}
	return dummy.Next
}

func getLen(l *ListNode) int {
	ans := 0
	for l != nil {
		ans++
		l = l.Next
	}
	return ans
}
