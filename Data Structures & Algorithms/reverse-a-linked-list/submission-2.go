/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // Recursion
 // it is easy to recognize the base cases are:
 // - if the link is empty, return the head
 // - if the link has one node, return the head
 // if the link has more than one node, we need to turn it into base case
 // for example, [0, 1, 2]
 // step 1. revert the rest of the list [1, 2] until reaching the tail, it is new head
 // step 2. change the direction: next = head.Next, next.Next = head, head=nil
 // return newHead

func reverseList(head *ListNode) *ListNode {
    // base cases
	// 1. edge case
	if head == nil {
		return nil
	}

	// 2. reach the tail
	if head.Next == nil {
		return head
	}

	// assume the rest of link are reversed
	newHead := reverseList(head.Next)

	// change direction logic
	// if next of node A is B, reverse the redirection will be
	head.Next.Next = head
	// break original link
	head.Next = nil

	return newHead
}
