package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Item struct {
	value    ListNode
	priority int
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index -= 1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value ListNode, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
