package main

import (
	"fmt"
	"week-3/sort"
)

func displayMergeSort() {
	l := []int{50, 10, 5, 15, 20, 0, -5, 30, -10}
	fmt.Printf("Data for mergesort: %v\n", l)

	sort.SortMerge(l)
	fmt.Printf("Data after mergesort: %v\n", l)
}

func displayMergeSortBU() {
	l := []int{50, 10, 5, 15, 20, 0, -5, 30, -10}
	fmt.Printf("Data for bottom-up mergesort: %v\n", l)

	sort.SortMerge(l)
	fmt.Printf("Data after bottom-up mergesort: %v\n", l)
}

func displayQuickSort() {
	l := []int{50, 10, 5, 15, 20, 0, -5, 30, -10}
	fmt.Printf("Data for quicksort: %v\n", l)

	sort.SortQuick(l)
	fmt.Printf("Data after quicksort: %v\n", l)
}

func main() {
	displayMergeSort()
	fmt.Println()

	displayMergeSortBU()
	fmt.Println()

	displayQuickSort()
	fmt.Println()
}
