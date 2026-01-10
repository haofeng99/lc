package linklist

// 1634 Add Two Polynomials Represented as Linked Lists
// 思路: 本质就是两个有序链表合并的变种
type PolyNode struct {
	Coefficient int
	Power       int
	Next        *PolyNode
}

func addPoly(poly1, poly2 *PolyNode) *PolyNode {
	dummy := &PolyNode{}
	cur := dummy
	for poly1.Next != nil && poly2.Next != nil {
		if poly1.Power > poly2.Power {
			cur.Next = &PolyNode{
				Coefficient: poly1.Coefficient,
				Power:       poly1.Power,
			}
			poly1 = poly1.Next
			cur = cur.Next
		} else if poly1.Power < poly2.Power {
			cur.Next = &PolyNode{
				Coefficient: poly2.Coefficient,
				Power:       poly2.Power,
			}
			poly2 = poly2.Next
			cur = cur.Next
		} else {
			sum := poly1.Coefficient + poly2.Coefficient
			if sum != 0 {
				cur.Next = &PolyNode{
					Coefficient: sum,
					Power:       poly1.Power,
				}
				cur = cur.Next
			}
			poly1 = poly1.Next
			poly2 = poly2.Next
		}
	}
	if poly1.Next != nil {
		cur.Next = poly1
	}
	if poly2.Next != nil {
		cur.Next = poly2
	}
	return dummy.Next
}
