package priorityqueue

import (
	"container/heap"
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

type IntervalHeap []*Interval

func (h *IntervalHeap) Len() int           { return len(*h) }
func (h *IntervalHeap) Less(i, j int) bool { return (*h)[i].End < (*h)[j].End }
func (h *IntervalHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntervalHeap) Push(x any) {
	*h = append(*h, x.(*Interval))
}

func (h *IntervalHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MinMeetingRooms(intervals []*Interval) int {

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	h := &IntervalHeap{intervals[0]}
	heap.Init(h)

	for i := 1; i < len(intervals); i++ {
		last := heap.Pop(h)
		if intervals[i].Start < last.(*Interval).End {
			heap.Push(h, last)
		}
		heap.Push(h, intervals[i])
	}

	return h.Len()
}

func Show_MinMeetingRooms() {
	intervals := []*Interval{{1, 5}, {7, 11}, {4, 8}, {9, 3}}
	res := MinMeetingRooms(intervals)
	fmt.Println(res)
}

// https://www.lintcode.com/problem/919/description
// solution PriorityQueue https://www.lintcode.com/problem/919/solution/57828

func MinMeetingRoomsII(intervals []*Interval) int {

	start, end := make([]int, len(intervals)), make([]int, len(intervals))

	for i, _ := range intervals {
		start[i] = intervals[i].Start
		end[i] = intervals[i].End
	}

	sort.Ints(start)
	sort.Ints(end)

	startIdx, endIdx := 0, 0
	res, count := 0, 0

	for startIdx < len(intervals) {
		if start[startIdx] < end[endIdx] {
			startIdx++
			count++
		} else {
			endIdx++
			count--
		}

		if count > res {
			res = count
		}
	}

	return res
}
