package iterator

type LinkedListIterator struct {
	current *Node
}

func (li *LinkedListIterator) HasNext() bool {
	return li.current != nil
}

func (li *LinkedListIterator) Next() interface{} {
	item := li.current.item
	li.current = li.current.next

	return item
}

func NewLinkedListIterator(l *LinkedList) Iterator {
	return &LinkedListIterator{
		current: l.first,
	}
}
