package presum

// https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
// 思路:前缀和相同的位置之间的元素和为0
// https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/solutions/46077/java-hashmap-liang-ci-bian-li-ji-ke-by-shane-34/
func removeZeroSumSublists(head *ListNode) *ListNode {
	preSumMap := make(map[int]*ListNode)
	dymmy := &ListNode{0, head}

	cur := dymmy
	preSum := 0
	for cur != nil {
		preSum += cur.Val
		preSumMap[preSum] = cur
		cur = cur.Next
	}

	cur = dymmy
	preSum = 0
	for cur != nil {
		preSum += cur.Val
		cur.Next = preSumMap[preSum].Next
		cur = cur.Next
	}
	return dymmy.Next
}

// 为什么需要一个dummy节点,以及为什么前缀和需要从dummy节点开始计算
// [1, -1] 最终结果需要时 []
