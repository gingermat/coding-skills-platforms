package main

import (
	"fmt"
	"math/rand"
	"time"
	"week-2/sort"
)

func prepareSlice(length int) []int {
	ti := make([]int, length)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		ti[i] = rand.Intn(999) - rand.Intn(999)
	}

	return ti
}

func displaySortSelection() {
	ti := prepareSlice(20)
	fmt.Printf("ti %v\n", ti)

	sort.SortSelection(ti)
	fmt.Printf("Returned after SortSelection: %v\n", ti)
}

func displaySortInsertion() {
	ti := prepareSlice(20)
	fmt.Printf("ti %v\n", ti)

	sort.SortInsertion(ti)
	fmt.Printf("Returned after SortInsertion: %v\n", ti)
}

func displaySortShell() {
	ti := prepareSlice(20)
	fmt.Printf("ti %v\n", ti)

	sort.SortShell(ti)
	fmt.Printf("Returned after SortShell: %v\n", ti)
}

func displayShuffle() {
	ti := prepareSlice(20)
	fmt.Printf("ti %v\n", ti)

	sort.Shuffle(ti)
	fmt.Printf("Returned after Shuffle: %v\n", ti)
}

func displayConvexHull() {
	a := sort.NewPoint2D(0, 0)
	b := sort.NewPoint2D(0, 1)
	c := sort.NewPoint2D(1, 0)

	fmt.Printf("Points: %v %v %v\n", a, b, c)

	fmt.Printf("Rotate clockwise [a->b->c]: %v\n", sort.CWW(a, b, c))
	fmt.Printf("Rotate counter-clockwise [c->b->a]: %v\n", sort.CWW(c, b, a))
	fmt.Printf("Rotate collinear [a->b->a]: %v\n", sort.CWW(a, b, a))
}

func main() {
	displaySortSelection()
	fmt.Println()

	displaySortInsertion()
	fmt.Println()

	displaySortShell()
	fmt.Println()

	displayShuffle()
	fmt.Println()

	displayConvexHull()
	fmt.Println()
}
