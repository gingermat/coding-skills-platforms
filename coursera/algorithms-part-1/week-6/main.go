package main

import (
	"fmt"
	"week-6/hash_table"
)

func displaySeparateChainingHashST() {
	st := hash_table.NewSeparateChainingHashST()
	fmt.Printf("st=%v\n", st)

	st.Put("Key1", "Value1")
	fmt.Printf("After st.Put(\"Key1\", \"Value1\"): %v\n", st)

	st.Put("Key2", "Value2")
	fmt.Printf("After st.Put(\"Key2\", \"Value2\"): %v\n", st)

	st.Put("Key1", "Value3")
	fmt.Printf("After update st.Put(\"Key1\", \"Value3\"): %v\n", st)

	v := st.Get("NonExists")
	fmt.Printf("After st.Get(\"NonExists\"): %v, st: %v\n", v, st)

	v = st.Get("Key1")
	fmt.Printf("After st.Get(\"Key1\"): %v, st: %v\n", v, st)
}

func displayLinearProbingHashST() {
	st := hash_table.NewLinearProbingHashST()
	fmt.Printf("st=%v\n", st)

	st.Put("Key1", "Value1")
	fmt.Printf("After st.Put(\"Key1\", \"Value1\"): %v\n", st)

	st.Put("Key2", "Value2")
	fmt.Printf("After st.Put(\"Key2\", \"Value2\"): %v\n", st)

	st.Put("Key1", "Value3")
	fmt.Printf("After update st.Put(\"Key1\", \"Value3\"): %v\n", st)

	v := st.Get("NonExists")
	fmt.Printf("After st.Get(\"NonExists\"): %v, st: %v\n", v, st)

	v = st.Get("Key1")
	fmt.Printf("After st.Get(\"Key1\"): %v, st: %v\n", v, st)
}

func main() {
	fmt.Printf("Display Separate Chaining Hash Symbol Table\n\n")

	displaySeparateChainingHashST()
	fmt.Println()

	fmt.Printf("Display Linear Probing Hash Symbol Table\n\n")

	displayLinearProbingHashST()
	fmt.Println()
}
