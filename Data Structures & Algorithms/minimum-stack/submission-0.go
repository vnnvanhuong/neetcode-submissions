// use a stack to maintain the elements
// also, use a stack with prefix-sum technique to maintain min value
type MinStack struct {
	elements []int
	minElements []int
}

func Constructor() MinStack {
	return MinStack {
		elements: []int{},
		minElements: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)

	m := val
	if len(this.minElements) > 0 {
		t := this.minElements[len(this.minElements)-1]
		if t < m {
			m = t
		}
	}

	this.minElements = append(this.minElements, m)
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:len(this.elements)-1]
	this.minElements = this.minElements[:len(this.minElements)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElements[len(this.minElements)-1]
}
