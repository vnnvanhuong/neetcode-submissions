type Node struct {
    val int
    next *Node
    prev *Node
}

type MyLinkedList struct {
    left *Node
    right *Node
}


func Constructor() MyLinkedList {
    l, r := &Node{}, &Node{}
    l.next = r
    r.prev = l

    return MyLinkedList{
        left: l,
        right: r,
    }
}


func (this *MyLinkedList) Get(index int) int {
    curr := this.left.next
    for index > 0 && curr != nil {
        curr = curr.next
        index -= 1
    }

    if index == 0 && curr != nil && curr != this.right {
        return curr.val
    }

    return -1
    
}


func (this *MyLinkedList) AddAtHead(val int)  {
    node, prev, next := &Node{val: val}, this.left, this.left.next

    node.next = next
    node.prev = prev
    prev.next = node
    next.prev = node
}


func (this *MyLinkedList) AddAtTail(val int)  {
    node, prev, next := &Node{val: val}, this.right.prev, this.right

    node.next = next
    node.prev = prev
    prev.next = node
    next.prev = node
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    curr := this.left.next;
    for index > 0 && curr != nil {
        curr = curr.next
        index -= 1
    }

    if index == 0 && curr != nil {
        node, prev, next := &Node{val: val}, curr.prev, curr

        node.next = next
        node.prev = prev
        prev.next = node
        next.prev = node
    }
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    curr := this.left.next;
    for index > 0 && curr != nil {
        curr = curr.next
        index -= 1
    }

    if index == 0 && curr != nil && curr != this.right {
        prev, next := curr.prev, curr.next

        prev.next = next
        next.prev = prev
    }
    
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
