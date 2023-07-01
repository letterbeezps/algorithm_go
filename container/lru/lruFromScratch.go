package lru

type LRUCacheFromScratch struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}
type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func ConstructorFromScratch(capacity int) LRUCacheFromScratch {
	return LRUCacheFromScratch{
		Keys: make(map[int]*Node),
		Cap:  capacity,
	}
}

func (cache *LRUCacheFromScratch) Add(node *Node) {
	node.Prev = nil
	node.Next = cache.head
	if cache.head != nil {
		cache.head.Prev = node
	}
	cache.head = node
	if cache.tail == nil {
		cache.tail = node
		cache.tail.Next = nil
	}
}

func (cache *LRUCacheFromScratch) Remove(node *Node) {
	if node == cache.head {
		cache.head = node.Next
		if cache.head != nil {
			cache.head.Prev = nil
		}
		node.Next = nil
		return
	}
	if node == cache.tail {
		cache.tail = node.Prev
		cache.tail.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (cache *LRUCacheFromScratch) Get(key int) int {
	if node, ok := cache.Keys[key]; ok {
		cache.Remove(node)
		cache.Add(node)
		return node.Val
	} else {
		return -1
	}
}

func (cache *LRUCacheFromScratch) Put(key int, value int) {
	if node, ok := cache.Keys[key]; ok {
		node.Val = value
		cache.Remove(node)
		cache.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		cache.Keys[key] = node
		cache.Add(node)
	}

	if len(cache.Keys) > cache.Cap {
		delete(cache.Keys, cache.tail.Key)
		cache.Remove(cache.tail)
	}
}
