package middle_of_ll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	middle := head

	for h, even := head, false; h != nil; h, even = h.Next, !even {
		if even {
			middle = middle.Next
		}
	}

	return middle
}
