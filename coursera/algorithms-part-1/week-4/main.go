package main

import (
	"fmt"
	"week-4/heap"
	"week-4/queue"
	"week-4/symbol_table"
)

func displayUnorderedMaxPQ() {
	pq := queue.NewUnorderedMaxPQ(10)
	fmt.Printf("pq=%v\n", pq)

	pq.Insert(15)
	fmt.Printf("After pq.Insert(15): %v\n", pq)

	pq.Insert(3)
	fmt.Printf("After pq.Insert(3): %v\n", pq)

	pq.Insert(12)
	fmt.Printf("After pq.Insert(12): %v\n", pq)

	pq.Insert(7)
	fmt.Printf("After pq.Insert(7): %v\n", pq)

	v := pq.DelMax()
	fmt.Printf("After pq.DelMax(): %v, value = %v\n", pq, v)

	v = pq.DelMax()
	fmt.Printf("After pq.DelMax(): %v, value = %v\n", pq, v)
}

func displayMaxPriorityQueue() {
	mq := queue.NewMaxPriorityQueue(10)
	fmt.Printf("mq=%v\n", mq)

	mq.Insert(15)
	fmt.Printf("After mq.Insert(15): %v\n", mq)

	mq.Insert(7)
	fmt.Printf("After mq.Insert(7): %v\n", mq)

	mq.Insert(2)
	fmt.Printf("After mq.Insert(2): %v\n", mq)

	mq.Insert(11)
	fmt.Printf("After mq.Insert(11): %v\n", mq)

	mq.Insert(4)
	fmt.Printf("After mq.Insert(4): %v\n", mq)

	mq.Insert(9)
	fmt.Printf("After mq.Insert(9): %v\n", mq)

	v := mq.DelMax()
	fmt.Printf("After mq.DelMax(): %v, value = %v\n", mq, v)

	v = mq.DelMax()
	fmt.Printf("After mqpq.DelMax(): %v, value = %v\n", mq, v)

	mq.Insert(4)
	fmt.Printf("After mq.Insert(4): %v\n", mq)

	mq.Insert(9)
	fmt.Printf("After mq.Insert(4): %v\n", mq)

	v = mq.DelMax()
	fmt.Printf("After mq.DelMax(): %v, value = %v\n", mq, v)

	v = mq.DelMax()
	fmt.Printf("After mq.DelMax(): %v, value = %v\n", mq, v)
}

func displayHeapSort() {
	l := []int{15, 3, 8, 2, 10, 22, 4, 6, 2}
	fmt.Printf("Data before heap sort=%v\n", l)

	heap.SortHeap(l)
	fmt.Printf("Data after heap sort: %v\n", l)
}

func displaySymbolTableLL() {
	var st symbol_table.SymbolTable = symbol_table.NewSymbolTableLL()
	fmt.Printf("st=%v\n", st)

	st.Put("Key1", "Value1")
	fmt.Printf("After st.Put(\"Key1\", \"Value1\"): %v\n", st)

	st.Put("Key2", "Value2")
	fmt.Printf("After st.Put(\"Key2\", \"Value2\"): %v\n", st)

	fmt.Printf("st.Contains(\"Key2\"): %v\n", st.Contains("Key2"))
	fmt.Printf("st.Contains(\"Key3\"): %v\n", st.Contains("Key3"))

	fmt.Printf("st.Keys(): %v\nst.Size(): %v\n", st.Keys(), st.Size())

	st.Delete("NonExists")
	fmt.Printf("After st.Delete(\"NonExists\"): %v, st: %v\n", st.Keys(), st)

	v := st.Get("NonExists")
	fmt.Printf("After st.Get(\"NonExists\"): %v, st: %v\n", v, st)

	v = st.Get("Key1")
	fmt.Printf("After st.Get(\"Key1\"): %v, st: %v\n", v, st)

	st.Delete("Key1")
	fmt.Printf("After st.Delete(\"Key1\"): %v, st: %v\n", st.Keys(), st)

	st.Put("Key3", "Value3")
	fmt.Printf("After st.Put(\"Key3\", \"Value3\"): %v\n", st)

	st.Delete("Key2")
	fmt.Printf("After st.Delete(\"Key2\"): %v, st: %v\n", st.Keys(), st)

	fmt.Printf("st.Size(): %v, st: %v\n", st.Size(), st)

	st.Delete("Key3")
	fmt.Printf("After st.Delete(\"Key3\"): %v, st: %v\n", st.Keys(), st)

	fmt.Printf("st.IsEmpty(): %v, st: %v\n", st.IsEmpty(), st)
}

func displaySymbolTableBST() {
	var st symbol_table.SymbolTable = symbol_table.NewSymbolTableBST()
	fmt.Printf("st=%v\n", st)

	st.Put("Key1", "Value1")
	fmt.Printf("After st.Put(\"Key1\", \"Value1\"): %v\n", st)

	st.Put("Key2", "Value2")
	fmt.Printf("After st.Put(\"Key2\", \"Value2\"): %v\n", st)

	fmt.Printf("st.Contains(\"Key2\"): %v\n", st.Contains("Key2"))
	fmt.Printf("st.Contains(\"Key3\"): %v\n", st.Contains("Key3"))

	fmt.Printf("st.Keys(): %v\nst.Size(): %v\n", st.Keys(), st.Size())

	st.Delete("NonExists")
	fmt.Printf("After st.Delete(\"NonExists\"): %v, st: %v\n", st.Keys(), st)

	v := st.Get("NonExists")
	fmt.Printf("After st.Get(\"NonExists\"): %v, st: %v\n", v, st)

	v = st.Get("Key1")
	fmt.Printf("After st.Get(\"Key1\"): %v, st: %v\n", v, st)

	st.Delete("Key1")
	fmt.Printf("After st.Delete(\"Key1\"): %v, st: %v\n", st.Keys(), st)

	st.Put("Key3", "Value3")
	fmt.Printf("After st.Put(\"Key3\", \"Value3\"): %v\n", st)

	st.Delete("Key2")
	fmt.Printf("After st.Delete(\"Key2\"): %v, st: %v\n", st.Keys(), st)

	fmt.Printf("st.Size(): %v, st: %v\n", st.Size(), st)

	st.Delete("Key3")
	fmt.Printf("After st.Delete(\"Key3\"): %v, st: %v\n", st.Keys(), st)

	fmt.Printf("st.IsEmpty(): %v, st: %v\n", st.IsEmpty(), st)
}

func main() {
	fmt.Printf("Display UnorderedMax Priority Queue\n\n")

	displayUnorderedMaxPQ()
	fmt.Println()

	fmt.Printf("Display Max Priority Queue\n\n")

	displayMaxPriorityQueue()
	fmt.Println()

	fmt.Printf("Display HeapSort\n\n")

	displayHeapSort()
	fmt.Println()

	fmt.Printf("Display Symbol Table(Linked List)\n\n")

	displaySymbolTableLL()
	fmt.Println()

	fmt.Printf("Display Symbol Table(Binary Search Tree)\n\n")

	displaySymbolTableBST()
	fmt.Println()
}
