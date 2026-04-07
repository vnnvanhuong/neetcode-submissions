/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // one naive approach is to copy value to array
 // then loop through items of array from n-1 to 0, and return the head

func reverseList(head *ListNode) *ListNode {
	vals := []int{}

	cur := head
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}

	var newHead *ListNode
	var newCur *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		node := &ListNode{
			Val: vals[i],
		}

		if newCur == nil {
			newCur = node
			newHead = newCur
			continue
		}

		newCur.Next = node
		newCur = newCur.Next
	}

	return newHead
    
}
