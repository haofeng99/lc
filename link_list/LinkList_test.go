package linklist

import (
	"testing"
)

func Test_sortOddEvent_splitList(t *testing.T) {
	h := CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	h1, h2 := _splitList(h)
	PrintListNode(h1)
	PrintListNode(h2)
}
