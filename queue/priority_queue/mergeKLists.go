package priorityqueue

import "container/heap"

type ListNodeQueue []*ListNode

func (h *ListNodeQueue) Len() int           { return len(*h) }
func (h *ListNodeQueue) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *ListNodeQueue) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *ListNodeQueue) Push(x any) {
	node := x.(*ListNode)
	*h = append(*h, node)
}

func (h *ListNodeQueue) Pop() any {
	n := len(*h)
	node := (*h)[n-1]
	*h = (*h)[:n-1]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	priority_queue := &ListNodeQueue{}
	heap.Init(priority_queue)

	for _, nodePtr := range lists {
		if nodePtr != nil {
			heap.Push(priority_queue, nodePtr)
		}
	}

	dummy := &ListNode{}
	cur := dummy
	for priority_queue.Len() > 0 {
		node := heap.Pop(priority_queue).(*ListNode)
		cur.Next = node
		if node.Next != nil {
			heap.Push(priority_queue, node.Next)
		}
		cur = node
	}

	return dummy.Next
}
