package linklist

import "fmt"

// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 只需判断快指针有没有走到头
func hasCycleII(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 进阶1 找出环的入口
// 方法1: 一次遍历,使用map缓存已经遍历的节点,当遇到重复节点时,即为环的入口
// map[ListNode]interface{}：这里的键是 ListNode 值类型。这意味着每次访问 map 时，都会复制整个 ListNode 结构体。如果 ListNode 结构体比较大，这会带来性能开销。此外，如果链表节点包含指针字段，比较两个 ListNode 值时会比较它们的内容，而不是它们的地址。
// map[*ListNode]interface{}：这里的键是 *ListNode 指针类型。这意味着 map 的键是节点的地址，而不是节点的内容。这样可以避免复制整个结构体，并且比较两个指针时只比较它们的地址，这样更高效。
func findCycleEntrance(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]interface{})
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = nil
		head = head.Next
	}
	return nil
}

// 方法2: 快慢指针,当快慢指针相遇时,让其中一个指针指向头节点,然后两个指针同时移动,再次相遇的节点即为环的入口
func findCycleEntranceII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	// 检测环的存在
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 如果没有环，返回 nil
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 找到环的入口
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

// 进阶2 环中节点的个数
// 方法1: 一次遍历,使用map缓存已经遍历的节点以及它们的位置,当遇到某个节点的下一个节点已经出现过时,
// 可以使用当前节点下标-下一个节点下标即为环中节点的个数
// 方法2: 快慢指针,当快慢指针相遇时,慢指针继续前进,快指针回到头节点,两个指针同时前进,再次相遇时,即为环中节点的个数
func countCycleNodeNum(head *ListNode) int {
	nodePosMap := make(map[*ListNode]int)
	curPos := 0
	for head != nil {
		if pos, ok := nodePosMap[head]; ok {
			return curPos - pos
		}
		nodePosMap[head] = curPos
		head = head.Next
		curPos++
	}
	return 0
}

//https://leetcode.cn/problems/linked-list-cycle/

func Show_hasCycle() {
	head := CreateList([]int{3, 2, 0, -4})
	head.Next.Next.Next.Next = head

	// entrenceNode := findCycleEntrance(head)
	// PrintNode(entrenceNode)

	countCycleNodeNum := countCycleNodeNum(head)
	fmt.Println(countCycleNodeNum)
}
