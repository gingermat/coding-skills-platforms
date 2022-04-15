package queue

type Node struct {
	item interface{}
	next *Node
}

type QueueOnLinkedList struct {
	first, last *Node
}

func (q *QueueOnLinkedList) Enqueue(item interface{}) {
	oldLast := q.last
	q.last = &Node{item: item}

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

func (q *QueueOnLinkedList) Dequeue() interface{} {
	item := q.first.item
	q.first = q.first.next

	if q.IsEmpty() {
		q.last = nil
	}

	return item
}

func (q *QueueOnLinkedList) IsEmpty() bool {
	return q.first == nil
}

func NewQueueOnLinkedList() *QueueOnLinkedList {
	return &QueueOnLinkedList{}
}
