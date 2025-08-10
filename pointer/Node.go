package pointer

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintNode(node *ListNode) {
	fmt.Println(node.Val)
}

func PrintListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func CreateList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}
