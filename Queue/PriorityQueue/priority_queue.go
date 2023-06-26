package Queue

import (
	"container/heap"
)

type Item[T any] struct {
	Value    T
	priority int
	index    int
}
type Queue[T any] []*Item[T]

func (q Queue[T]) Len() int {
	return len(q)
}

func (q Queue[T]) Less(i, j int) bool {
	return q[i].priority > q[j].priority
}

func (q Queue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = j
	q[j].index = i
}

func (q *Queue[T]) Push(x any) {
	n := len(*q)
	item := x.(*Item[T])
	item.index = n
	*q = append(*q, item)
}

func (q *Queue[T]) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index -= 1
	*q = old[0 : n-1]
	return item
}

func (q *Queue[T]) update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(q, item.index)
}

type PriorityQueue[T any] struct {
	Queue Queue[T]
}

func (q *PriorityQueue[T]) Init() {
	heap.Init(&q.Queue)
}
func (q *PriorityQueue[T]) Len() int {
	return q.Queue.Len()
}

func (q *PriorityQueue[T]) Pop() any {
	return heap.Pop(&q.Queue)
}

func (q *PriorityQueue[T]) Push(x any) {
	heap.Push(&q.Queue, x)
}
