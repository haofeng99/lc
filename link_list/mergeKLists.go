package linklist

import "container/heap"

// https://leetcode.cn/problems/merge-k-sorted-lists/description

type ListNodeHeap []*ListNode

func (h *ListNodeHeap) Len() int           { return len(*h) }
func (h *ListNodeHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *ListNodeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	priority_queue := &ListNodeHeap{}
	heap.Init(priority_queue)

	// 利用链表有序的特点, 只把链表的头节点放入优先级队列
	for i := 0; i < len(lists); i++ {
		head := lists[i]
		if head != nil {
			heap.Push(priority_queue, head)
		}
	}

	dummy := &ListNode{}
	cur := dummy

	for priority_queue.Len() > 0 {
		popNode := heap.Pop(priority_queue).(*ListNode)
		cur.Next = popNode
		cur = cur.Next
		if popNode.Next != nil {
			heap.Push(priority_queue, popNode.Next)
		}
	}
	return dummy.Next
}
