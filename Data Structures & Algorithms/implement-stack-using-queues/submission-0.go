type Node struct {
	val int
	next *Node
	prev *Node
}

type MyQueue struct {
	left *Node
	right *Node
	size int
}

func NewMyQueue() MyQueue {
	left := &Node{}
	right := &Node{}

	left.next = right
	right.prev = left

	return MyQueue{
		left: left,
		right: right,
	}
} 

func (this *MyQueue) Push(x int) {
	node, prev, next := &Node{val: x}, this.right.prev, this.right

	node.next = next
	node.prev = prev
	prev.next = node
	next.prev = node

	this.size += 1
}

func (this *MyQueue) Peek() int {
	// empty queue, return 0
	return this.left.next.val
}

func (this *MyQueue) Pop() int {
	node, prev, next := this.left.next, this.left, this.left.next.next

	prev.next = next
	next.prev = prev

	this.size -= 1

	return node.val
}

func (this *MyQueue) Empty() bool {
	return this.left.next == this.right
}

func (this *MyQueue) Size() int {
	return this.size
}

type MyStack struct {
	queue MyQueue
}

func Constructor() MyStack {
	queue := NewMyQueue()
	return MyStack{queue: queue}
}

// [1 2 3 4]
// [2 3 4 1]

func (this *MyStack) Push(x int) {
	rotates := this.queue.Size()

	this.queue.Push(x)

	// rotate elements
	for rotates > 0 {
		top := this.queue.Pop()
		this.queue.Push(top)
		rotates--
	}
}

func (this *MyStack) Pop() int {
	return this.queue.Pop()

}

func (this *MyStack) Top() int {
	return this.queue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
