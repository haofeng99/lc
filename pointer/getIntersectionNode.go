package pointer

// https://leetcode.cn/problems/intersection-of-two-linked-lists/

// a
//
//		\
//		 \
//		  \
//		    c - - - - - - - - d
//		  /
//		 /
//	  b
//
// 结论: ac + cd + bc == bc + cd + ac
// 如果两个链表有相交点,按照上面结论的路径走,则两个指针会汇聚在c点,否则会最终指向nil
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
