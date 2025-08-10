package linklist

// https://leetcode.cn/problems/add-two-numbers-ii/
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	flag := 0
	idx1, idx2 := len(stack1)-1, len(stack2)-1
	var node *ListNode
	for idx1 >= 0 || idx2 >= 0 {
		v1 := 0
		if idx1 >= 0 {
			v1 = stack1[idx1]
		}

		v2 := 0
		if idx2 >= 0 {
			v2 = stack2[idx2]
		}

		sum := flag + v1 + v2
		nv := sum % 10
		flag = sum / 10

		curNode := &ListNode{nv, node}
		node = curNode

		idx1--
		idx2--
	}
	if flag != 0 {
		curNode := &ListNode{flag, node}
		node = curNode
	}
	return node
}
