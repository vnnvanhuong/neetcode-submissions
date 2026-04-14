type Node struct {
    url string
    next *Node
    prev *Node 
}

type BrowserHistory struct {
    left *Node
    right *Node
    curr *Node
}


func Constructor(homepage string) BrowserHistory {
    node := &Node{url: homepage}
    left := &Node{}
    right:= &Node{}

    node.next = right
    node.prev = left
    left.next = node
    right.prev = node
    return BrowserHistory{
        left: left,
        right: right,
        curr: node,
    }
}

func (this *BrowserHistory) Visit(url string)  {
    node, prev, next := &Node{url: url}, this.curr, this.right

    node.next = next
    node.prev = prev
    prev.next = node
    next.prev = node

    this.curr = this.curr.next
}


func (this *BrowserHistory) Back(steps int) string {
    for steps > 0 && this.curr != this.left.next {
        this.curr = this.curr.prev
        steps -= 1
    }

    if this.curr != this.left.next && steps == 0 {
        return this.curr.url
    }

    return this.left.next.url
    
}


func (this *BrowserHistory) Forward(steps int) string {
    for this.curr != this.right.prev && steps > 0 {
        this.curr = this.curr.next
        steps -= 1
    }

    if this.curr != this.right.prev && steps == 0 {
        return this.curr.url
    }

    return this.right.prev.url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
