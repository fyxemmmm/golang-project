package favourite

type LRUCache struct {
	capacity int
	size int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct{
	key, value int
	prev, next *DLinkedNode
}


func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache: map[int]*DLinkedNode{},
		head: &DLinkedNode{},
		tail: &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := &DLinkedNode{
			key: key,
			value: value,
		}
		this.size ++
		this.cache[key] = node
		this.addToHead(node)
		if this.size > this.capacity {
			this.size --
			removed := this.removeTail()
			delete(this.cache, removed.key)
		}
	}else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}


func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}


func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

