package remove_nth_node

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		nodeCount        int
		beforeDeleteNode *ListNode
	)

	for h := head; h != nil; h = h.Next {
		if nodeCount == n {
			beforeDeleteNode = head
		} else if nodeCount > n {
			beforeDeleteNode = beforeDeleteNode.Next
		}

		nodeCount++
	}

	if nodeCount == n {
		head = head.Next
	}

	if beforeDeleteNode != nil {
		beforeDeleteNode.Next = beforeDeleteNode.Next.Next
	}

	return head
}
