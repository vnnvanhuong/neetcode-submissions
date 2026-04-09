/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

    dummy := &ListNode{}
	l1, l2, d := list1, list2, dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			d.Next = l1
			l1 = l1.Next
			d = d.Next
			continue
		}

		d.Next = l2
		l2 = l2.Next
		d = d.Next
	}

	if l1 != nil {
		d.Next = l1
	}

	if l2 != nil {
		d.Next = l2
	}

	return dummy.Next
}
