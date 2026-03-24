// LeetCode L146

package main

func main() {

}

// circular doubly-linked list with a dummy head

type myListNode struct {
	Key, Value int
	Pre, Next  *myListNode
}

type LRUCache struct {
	head             *myListNode
	mp               map[int]*myListNode
	capacity, length int
}

func (lru *LRUCache) removeNode(node *myListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (lru *LRUCache) addNode(node *myListNode) {
	node.Next = lru.head.Next
	node.Pre = lru.head
	lru.head.Next = node
	node.Next.Pre = node
}

func Constructor(capacity int) LRUCache {
	head := &myListNode{}
	head.Pre = head
	head.Next = head
	return LRUCache{
		head:     head,
		mp:       make(map[int]*myListNode, 0),
		capacity: capacity,
		length:   0,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.mp[key]; ok {
		lru.removeNode(node)
		lru.addNode(node)
		return node.Value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.mp[key]; ok {
		lru.removeNode(node)
		node.Value = value
		lru.addNode(node)
	} else {
		if lru.length >= lru.capacity {
			delete(lru.mp, lru.head.Pre.Key)
			lru.removeNode(lru.head.Pre)
			lru.length--
		}
		node = &myListNode{
			Key:   key,
			Value: value,
		}
		lru.addNode(node)
		lru.mp[key] = node
		lru.length++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
