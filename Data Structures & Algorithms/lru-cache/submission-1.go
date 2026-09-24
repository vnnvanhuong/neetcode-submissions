type entry struct {
	key int
	val int
}

type LRUCache struct {
	cap int
	hm map[int]*list.Element
	list *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache {
		cap: capacity,
		hm: make(map[int]*list.Element, capacity),
		list: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.hm[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(entry).val
	}
    return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.hm[key]; ok {
		elem.Value = entry{key: key, val: value}
		this.list.MoveToFront(elem)
		return
	}
    if len(this.hm) >= this.cap {
		lru := this.list.Back()
		this.list.Remove(lru)
		delete(this.hm, lru.Value.(entry).key)
	}

	elem := this.list.PushFront(entry{key: key, val: value})
	this.hm[key] = elem
}
