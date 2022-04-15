package iterator

type Node struct {
	item interface{}
	next *Node
}

type LinkedList struct {
	first *Node
}

func NewLinkedList(items []interface{}) *LinkedList {
	var (
		first *Node
		last  *Node
	)

	for _, el := range items {
		oldlast := last
		last = &Node{item: el}

		if first == nil {
			first = last
		} else {
			oldlast.next = last
		}
	}

	return &LinkedList{
		first: first,
	}
}
