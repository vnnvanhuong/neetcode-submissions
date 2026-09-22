type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func findKthLargest(nums []int, k int) int {
	// init the min-heap from nums
	// when heap size < k, push to heap
	// when heap size = k, if nums[i] > heap[0], pop and push new element
	// return heap[0]
	h := &MinHeap{}
	heap.Init(h)
	for _, x := range nums {
		if h.Len() < k {
			heap.Push(h, x)
			continue
		}

		if x > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, x)
		} 
	}

	return heap.Pop(h).(int)
}
