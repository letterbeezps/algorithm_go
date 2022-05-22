package main

import "container/list"

type pair struct {
	k, v int
}

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.List.MoveToFront(node)
		return node.Value.(pair).v
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Value = pair{k: key, v: value}
		this.List.MoveToFront(node)
	} else {
		node := this.List.PushFront(pair{k: key, v: value})
		this.Keys[key] = node
	}

	if this.List.Len() > this.Cap {
		node := this.List.Back()
		this.List.Remove(node)
		delete(this.Keys, node.Value.(pair).k)
	}
}
