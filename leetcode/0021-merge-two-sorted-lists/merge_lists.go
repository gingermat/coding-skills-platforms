package merge_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var result = &ListNode{}
	tail := result

	for {
		if list1 == nil {
			tail.Next = list2
			break
		}

		if list2 == nil {
			tail.Next = list1
			break
		}

		if list1.Val >= list2.Val {
			tail.Next = list2
			list2 = list2.Next
		} else {
			tail.Next = list1
			list1 = list1.Next
		}

		tail = tail.Next
	}

	return result.Next
}
