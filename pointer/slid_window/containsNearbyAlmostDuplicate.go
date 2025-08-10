package slidwindow

import (
	"fmt"
	"math/rand"
)

type node struct {
	ch       [2]*node
	priority int
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val}
	}
	d := o.cmp(val)
	o.ch[d] = t._put(o.ch[d], val)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int) {
	t.root = t._delete(t.root, val)
}

func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c > 0:
			o = o.ch[1]
		default:
			return o
		}
	}
	return
}

// 滑动窗口+treeMap
// golang没有语言层面的实现,需要自己实现
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	set := &treap{}
	for i, v := range nums {
		if lb := set.lowerBound(v - t); lb != nil && lb.val <= v+t {
			return true
		}
		set.put(v)
		if i >= k {
			set.delete(nums[i-k])
		}
	}
	return false
}

// https://leetcode.cn/problems/contains-duplicate-iii/
// https://leetcode.cn/problems/contains-duplicate-iii/solutions/726905/gong-shui-san-xie-yi-ti-shuang-jie-hua-d-dlnv/
// 需要注意的是,为什么最后移除treeMap的元素事,当下标的等于k时就可以移除最开始放入的元素
// 因为对于下一个下标的元素来说, 最开始放入的元素已经超出了k的管辖范围,不会对最终结果产生作用,因此可以直接移除

func Show_containsNearbyAlmostDuplicate() {
	nums := []int{7, 1, 3}
	ans := containsNearbyAlmostDuplicate(nums, 2, 3)
	fmt.Println(ans)
}
