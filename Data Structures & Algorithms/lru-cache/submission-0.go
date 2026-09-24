type Node struct {
	prev *Node
	next *Node
	key int
	val int
}

type LRUCache struct {
	cap int
	hm map[int]*Node
	right *Node
	left *Node
}

func Constructor(capacity int) LRUCache {
	left, right := &Node{}, &Node{}
	left.next = right
	right.prev = left

	return LRUCache {
		cap: capacity,
		hm: make(map[int]*Node, capacity),
		left: left,
		right: right,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.hm[key]; exists {
		this.remove(node)
		this.insert(node)
		return node.val
	}
    return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.hm[key]; exists {
		node.val = value
		this.remove(node)
		this.insert(node)
		return
	}
    if len(this.hm) >= this.cap {
		lru := this.left.next
		this.remove(lru)
		delete(this.hm, lru.key)
	}

	node := &Node{key: key, val: value}
	this.insert(node)
	this.hm[key] = node
}

func (this *LRUCache) insert(node *Node) {
	prev, next := this.right.prev, this.right
	node.prev = prev
	node.next = next
	prev.next = node
	next.prev = node
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}
