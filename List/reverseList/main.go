package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(pre *ListNode, k int) *ListNode {
	last := pre
	for i := 0; i <= k; i++ {
		last = last.Next
		if i != k && last == nil {
			return nil
		}
	}
	// 将 pre 和 last之间的节点翻转
	tail := pre.Next
	curr := pre.Next.Next
	for curr != last {
		curr_next := curr.Next
		tail.Next = curr_next

		curr.Next = pre.Next
		pre.Next = curr
		curr = curr_next
	}
	return tail
}
