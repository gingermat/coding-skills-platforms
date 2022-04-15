package main

import (
	"fmt"
	"week-2/iterator"
	"week-2/queue"
	"week-2/stack"
)

func display_stack_linked_list() {
	st := stack.NewStackOnLinkedList()
	fmt.Printf("Stack st=%v\n", st)

	fmt.Printf("Check empty stack st.IsEmpty(): %v\n", st.IsEmpty())

	st.Push(1)
	fmt.Printf("Check empty stack st.IsEmpty() after st.Push(1): %v\n", st.IsEmpty())

	item := st.Pop()
	fmt.Printf("Returned element from stack st.Pop(): %v\n", item)

	fmt.Printf("Check empty stack st.IsEmpty() after st.Pop(): %v\n", st.IsEmpty())
}

func display_stack_array() {
	st := stack.NewStackOnArray()
	fmt.Printf("Stack st=%v\n", st)

	fmt.Printf("Check empty stack st.IsEmpty(): %v\n", st.IsEmpty())

	st.Push(11)
	fmt.Printf("Add value into stack st.Push(11): %v\n", st)

	st.Push(22)
	fmt.Printf("Add value into stack st.Push(22): %v\n", st)

	st.Push(33)
	fmt.Printf("Add value into stack st.Push(33): %v\n", st)

	st.Push(44)
	fmt.Printf("Add value into stack st.Push(44): %v\n", st)

	st.Push(55)
	fmt.Printf("Add value into stack st.Push(55): %v\n", st)

	fmt.Printf("Check empty stack st.IsEmpty(): %v\n", st.IsEmpty())

	fmt.Printf("Returned element from stack st.Pop(): %v, st: %v\n", st.Pop(), st)
	fmt.Printf("Returned element from stack st.Pop(): %v, st: %v\n", st.Pop(), st)
	fmt.Printf("Returned element from stack st.Pop(): %v, st: %v\n", st.Pop(), st)
	fmt.Printf("Returned element from stack st.Pop(): %v, st: %v\n", st.Pop(), st)
	fmt.Printf("Returned element from stack st.Pop(): %v, st: %v\n", st.Pop(), st)

	fmt.Printf("Check empty stack st.IsEmpty(): %v\n", st.IsEmpty())
}

func display_queue_linked_list() {
	q := queue.NewQueueOnLinkedList()
	fmt.Printf("Queue q=%v\n", q)

	fmt.Printf("Check empty queue q.IsEmpty(): %v\n", q.IsEmpty())

	q.Enqueue("a")
	fmt.Printf("Add value into queue q.Enqueue(\"a\"): %v\n", q)

	fmt.Printf("Check empty queue q.IsEmpty(): %v\n", q.IsEmpty())

	q.Enqueue(2)
	fmt.Printf("Add value into queue q.Enqueue(2): %v\n", q)

	fmt.Printf("Returned element from queue q.Dequeue(): %v, q: %v\n", q.Dequeue(), q)
	fmt.Printf("Returned element from queue q.Dequeue(): %v, q: %v\n", q.Dequeue(), q)

	fmt.Printf("Check empty queue q.IsEmpty(): %v\n", q.IsEmpty())
}

func display_iterator_linked_list() {
	ll := iterator.NewLinkedList([]interface{}{1, 2, 4, 8, 16})
	fmt.Printf("LinkedList ll=%v\n", ll)

	it := iterator.NewLinkedListIterator(ll)
	fmt.Printf("LinkedList Iterator it=%v\n", it)

	for it.HasNext() {
		fmt.Printf("Get item from LinkedList Iterator it.Next()=%v\n", it.Next())
	}
}

func main() {
	fmt.Printf("Display Stack(linked list)\n\n")

	display_stack_linked_list()
	fmt.Println()

	fmt.Printf("Display Stack(array)\n\n")

	display_stack_array()
	fmt.Println()

	fmt.Printf("Display Queue(linked list)\n\n")

	display_queue_linked_list()
	fmt.Println()

	fmt.Printf("Display Iterator(linked list)\n\n")

	display_iterator_linked_list()
	fmt.Println()
}
