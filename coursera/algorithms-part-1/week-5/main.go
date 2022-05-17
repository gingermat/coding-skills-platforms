package main

import (
	"fmt"
	"week-5/rb_tree"
)

func displayRedBlackBST() {
	var bst = rb_tree.NewRBBinarySearchTree()
	fmt.Printf("bst=%v\n", bst)

	bst.Put("Key1", "Value1")
	fmt.Printf("After bst.Put(\"Key1\", \"Value1\"): %v\n", bst)

	bst.Put("Key2", "Value2")
	fmt.Printf("After bst.Put(\"Key2\", \"Value2\"): %v\n", bst)

	v := bst.Get("NonExists")
	fmt.Printf("After bst.Get(\"NonExists\"): %v, bst: %v\n", v, bst)

	v = bst.Get("Key1")
	fmt.Printf("After bst.Get(\"Key1\"): %v, st: %v\n", v, bst)
}

func main() {
	fmt.Printf("Display Red Black Binary Search Tree\n\n")

	displayRedBlackBST()
	fmt.Println()
}
