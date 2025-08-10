package linklist

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

var listLen int

func Constructor() MyLinkedList {
	listLen = 0
	return MyLinkedList{-1, nil}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= listLen {
		return -1
	}
	idx := 0
	cur := this.next
	for idx < index && cur != nil {
		cur = cur.next
		idx++
	}
	if cur != nil {
		return cur.val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &MyLinkedList{val, nil}
	head := this.next
	this.next = newHead
	newHead.next = head
	listLen++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &MyLinkedList{val, nil}
	listLen++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > listLen {
		return
	}
	node := &MyLinkedList{val, nil}
	idx := 0
	pre := this
	for idx < index && pre != nil {
		pre = pre.next
		idx++
	}

	if pre == nil {
		return
	}

	pn := pre.next

	pre.next = node
	node.next = pn
	listLen++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= listLen {
		return
	}
	idx := 0
	pre := this
	for idx < index {
		pre = pre.next
		idx++
	}
	if pre == nil {
		return
	}
	pn := pre.next
	if pn != nil {
		pre.next = pn.next
	} else {
		pre.next = nil
	}
	listLen++
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
