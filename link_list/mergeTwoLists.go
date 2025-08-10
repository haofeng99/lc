package linklist

import "fmt"

// 思路
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	pre := &ListNode{}
	ans := pre
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 == nil {
		pre.Next = list2
	}
	if list2 == nil {
		pre.Next = list1
	}
	return ans.Next
}

// 递归方法
func mergeTwoListsRecure(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecure(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecure(list1, list2.Next)
		return list2
	}
}

// 进阶 需要考虑重复节点的去重问题
// 思路: 需要考虑三个地方的去重复
// 1.单个链表相邻节点的重复  2 -> 3 -> 3 -> 5
// 2.两个链表的节点重复
// 1 -> 3 -> 4 -> 6
// 2 -> 3 -> 4 -> 5
// 3.一个链表走完后,另一个链表还有重复节点
func mergeTwoListsNoDupNode(list1 *ListNode, list2 *ListNode) *ListNode {
	pre := &ListNode{-999, nil}
	ans := pre
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if pre.Val == list1.Val {
				list1 = list1.Next
				continue
			}
			pre.Next = list1
		} else if list1.Val > list2.Val {
			if pre.Val == list2.Val {
				list2 = list2.Next
				continue
			}
			pre.Next = list2
		} else {
			// 两个链表的节点重复
			list1 = list1.Next
			continue
		}
		pre = pre.Next
	}
	if list1 == nil {
		for list2 != nil {
			if list2.Next != nil && list2.Val == list2.Next.Val {
				list2 = list2.Next
				continue
			}
			if pre.Val == list2.Val {
				list2 = list2.Next
				continue
			}
			pre.Next = list2
			pre = pre.Next
		}
	}
	if list2 == nil {
		for list1 != nil {
			if list1.Next != nil && list1.Val == list1.Next.Val {
				list1 = list1.Next
				continue
			}
			if pre.Val == list1.Val {
				list1 = list1.Next
				continue
			}
			pre.Next = list1
			pre = pre.Next
		}
	}
	return ans.Next
}

// 递归方法
func mergeTwoListsNoDupRecure(list1 *ListNode, list2 *ListNode) *ListNode {

	for list1 != nil && list1.Next != nil && list1.Val == list1.Next.Val {
		list1 = list1.Next
	}
	for list2 != nil && list2.Next != nil && list2.Val == list2.Next.Val {
		list2 = list2.Next
	}

	if list1 == nil {
		for list2 != nil && list2.Next != nil && list2.Val == list2.Next.Val {
			list2 = list2.Next
		}
		if list2 != nil {
			list2.Next = mergeTwoListsNoDupRecure(list1, list2.Next)
		}
		return list2

	}
	if list2 == nil {
		for list1 != nil && list1.Next != nil && list1.Val == list1.Next.Val {
			list1 = list1.Next
		}
		if list1 != nil {
			list1.Next = mergeTwoListsNoDupRecure(list1.Next, list2)
		}
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsNoDupRecure(list1.Next, list2)
		return list1
	} else if list1.Val > list2.Val {
		list2.Next = mergeTwoListsNoDupRecure(list1, list2.Next)
		return list2
	} else {
		return mergeTwoListsNoDupRecure(list1.Next, list2)
	}
}

// https://leetcode.cn/problems/merge-two-sorted-lists/description/
func Show_MergeTwoLists() {
	pre := &ListNode{5, nil}
	fmt.Println(pre.Val)

	list1 := CreateList([]int{2, 3, 3, 5})
	list2 := CreateList([]int{})
	head := mergeTwoListsNoDupNode(list1, list2)
	PrintListNode(head)

	list3 := CreateList([]int{2, 3, 3, 5})
	list4 := CreateList([]int{1})
	head1 := mergeTwoListsNoDupRecure(list3, list4)
	PrintListNode(head1)
}
