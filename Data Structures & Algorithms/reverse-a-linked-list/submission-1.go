/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // three pointers
 // prev, curr, and next
 // inside the loop, save curr.Next,
 // curr.Next = prev
 // move prev, curr one step forward

func reverseList(head *ListNode) *ListNode {
   var prev, next *ListNode
   var curr = head
   if curr == nil {
	return nil 
   }

   for curr != nil {
	// save forward link
	next = curr.Next

	// reverse the link
	curr.Next = prev

	// move forward
	prev = curr
	curr = next
   }

   return prev
}


