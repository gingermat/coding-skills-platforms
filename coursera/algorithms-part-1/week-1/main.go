package main

import (
	"fmt"

	"week-1/quick_find"
	"week-1/quick_union"
	"week-1/quick_union_weighted"
)

func display_quick_find() {
	qf := quick_find.NewQF(10)
	fmt.Printf("qf=%v\n", qf)

	fmt.Println()

	qf.Union(5, 2)
	fmt.Printf("After qf.Union(5, 2): %v\n", qf)

	qf.Union(5, 3)
	fmt.Printf("After qf.Union(5, 3): %v\n", qf)

	qf.Union(1, 3)
	fmt.Printf("After qf.Union(1, 3): %v\n", qf)

	fmt.Println()

	fmt.Printf("Check connection qf.IsConnected(5, 2): %v\n", qf.IsConnected(5, 2))
	fmt.Printf("Check connection qf.IsConnected(2, 3): %v\n", qf.IsConnected(2, 3))
	fmt.Printf("Check connection qf.IsConnected(1, 3): %v\n", qf.IsConnected(1, 3))
}

func display_quick_union() {
	qu := quick_union.NewQU(10)
	fmt.Printf("qu=%v\n", qu)

	fmt.Println()

	qu.Union(5, 2)
	fmt.Printf("After qu.Union(5, 2): %v\n", qu)

	qu.Union(5, 3)
	fmt.Printf("After qu.Union(5, 3): %v\n", qu)

	qu.Union(1, 3)
	fmt.Printf("After qu.Union(1, 3): %v\n", qu)

	fmt.Println()

	fmt.Printf("Check connection qu.IsConnected(5, 2): %v\n", qu.IsConnected(5, 2))
	fmt.Printf("Check connection qu.IsConnected(2, 3): %v\n", qu.IsConnected(2, 3))
	fmt.Printf("Check connection qu.IsConnected(1, 3): %v\n", qu.IsConnected(1, 3))
}

func display_quick_union_weighted() {
	wqu := quick_union_weighted.NewWeightedQU(10)
	fmt.Printf("wqu=%v\n", wqu)

	fmt.Println()

	wqu.Union(5, 2)
	fmt.Printf("After wqu.Union(5, 2): %v\n", wqu)

	wqu.Union(5, 3)
	fmt.Printf("After wqu.Union(5, 3): %v\n", wqu)

	wqu.Union(1, 3)
	fmt.Printf("After wqu.Union(1, 3): %v\n", wqu)

	fmt.Println()

	fmt.Printf("Check connection wqu.IsConnected(5, 2): %v\n", wqu.IsConnected(5, 2))
	fmt.Printf("Check connection wqu.IsConnected(2, 3): %v\n", wqu.IsConnected(2, 3))
	fmt.Printf("Check connection wqu.IsConnected(1, 3): %v\n", wqu.IsConnected(1, 3))
}

func main() {
	fmt.Printf("Display Quick Find\n\n")

	display_quick_find()
	fmt.Println()

	fmt.Printf("Display Quick Union\n\n")

	display_quick_union()
	fmt.Println()

	fmt.Printf("Display Quick Union Weighted\n\n")

	display_quick_union_weighted()
	fmt.Println()
}
