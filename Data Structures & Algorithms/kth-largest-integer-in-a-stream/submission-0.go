type KthLargest struct {
	minHeap []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k:k}
	kl.minHeap = append(kl.minHeap, nums...)
	heap.Init(&kl)
	for len(kl.minHeap) > k {
		heap.Pop(&kl)
	}

	return kl
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this, val)
	if len(this.minHeap) > this.k {
		heap.Pop(this)
	}

	return this.minHeap[0]
}

func (h KthLargest) Len() int {
	return len(h.minHeap)
}

func (h KthLargest) Less(i, j int) bool {
	return h.minHeap[i] < h.minHeap[j]
}

func (h KthLargest) Swap(i, j int) {
	h.minHeap[i], h.minHeap[j] = h.minHeap[j], h.minHeap[i]
}

func (h *KthLargest) Push(x interface{}) {
	h.minHeap = append(h.minHeap, x.(int))
}

func (h *KthLargest) Pop() interface{} {
	x := h.minHeap[len(h.minHeap)-1]
	h.minHeap = h.minHeap[:len(h.minHeap)-1]
	return x
}
