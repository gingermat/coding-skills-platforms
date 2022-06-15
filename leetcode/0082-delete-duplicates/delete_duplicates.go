package delete_duplicates

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head = &ListNode{
		Next: head,
	}

	for n := head; n.Next != nil && n.Next.Next != nil; {
		if n.Next.Val != n.Next.Next.Val {
			n = n.Next
			continue
		}

		value := n.Next.Val

		for n.Next != nil && n.Next.Val == value {
			n.Next = n.Next.Next
		}
	}

	return head.Next
}
