type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	// insert all the weghts into the MaxHeap
	h := (*MaxHeap)(&stones)
	heap.Init(h)

	// while more than one element remains.
	// pop two elements
	// calculate the different d
	// if d == 0, continue
	// else push d into the heap
	for len(*h) > 1 {
		first := heap.Pop(h)
		second := heap.Pop(h)
		difference := first.(int) - second.(int)
		if difference == 0 {
			continue
		}

		heap.Push(h, difference)
	}

	if len(*h) == 1 {
		return (*h)[0]
	}

	return 0
}

