package lru

import "container/list"

type pair[K comparable, V any] struct {
	Key   K
	Value V
}

type LRUCache[K comparable, V any] struct {
	Cap  int
	Keys map[K]*list.Element
	List *list.List
}

func Constructor[K comparable, V any](capacity int) LRUCache[K, V] {
	return LRUCache[K, V]{
		Cap:  capacity,
		Keys: make(map[K]*list.Element),
		List: list.New(),
	}
}

func (this *LRUCache[K, V]) Get(key K) V {
	if node, ok := this.Keys[key]; ok {
		this.List.MoveToFront(node)
		return node.Value.(pair[K, V]).Value
	}

	var zero V
	return zero
}

func (this *LRUCache[K, V]) Put(key K, value V) {
	if node, ok := this.Keys[key]; ok {
		node.Value = pair[K, V]{Key: key, Value: value}
		this.List.MoveToFront(node)
	} else {
		node := this.List.PushFront(pair[K, V]{Key: key, Value: value})
		this.Keys[key] = node
	}

	if this.List.Len() > this.Cap {
		node := this.List.Back()
		this.List.Remove(node)
		delete(this.Keys, node.Value.(pair[K, V]).Key)
	}
}
