package Queue

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l < 1 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	tl := 0
	for _, v := range lists {
		if v != nil {
			tl++
		}
	}
	if tl == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	q := Queue[*ListNode]{}
	index := 0
	for _, v := range lists {
		if v != nil {
			q = append(q, &Item[*ListNode]{
				Value:    v,
				priority: v.Val,
				index:    index,
			})
			index++
		}
	}
	pq := PriorityQueue[*ListNode]{
		Queue: q,
	}
	pq.Init()
	for pq.Len() > 0 {
		item := pq.Pop().(*Item[*ListNode])
		node := item.Value
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			pq.Push(&Item[*ListNode]{
				Value:    node.Next,
				priority: node.Next.Val,
			})
		}
	}
	return dummy.Next
}
