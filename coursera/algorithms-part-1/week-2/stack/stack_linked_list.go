package stack

type Node struct {
	item interface{}
	next *Node
}

type StackOnLinkedList struct {
	first *Node
}

func (s *StackOnLinkedList) Push(item interface{}) {
	node := Node{item: item, next: s.first}
	s.first = &node
}

func (s *StackOnLinkedList) Pop() interface{} {
	node := s.first
	s.first = node.next

	return node.item
}

func (s *StackOnLinkedList) IsEmpty() bool {
	return s.first == nil
}

func NewStackOnLinkedList() *StackOnLinkedList {
	return &StackOnLinkedList{}
}
