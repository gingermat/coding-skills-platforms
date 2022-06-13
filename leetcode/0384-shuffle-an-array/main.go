package main

import (
	"math/rand"
	"reflect"
	"time"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())

	res := make([]int, len(this.nums))
	copy(res, this.nums)

	for i := len(res) - 1; i > 0; i-- {
		idx := rand.Intn(i + 1)
		res[idx], res[i] = res[i], res[idx]
	}
	return res
}

func main() {
	input := []int{1, 2, 3, 4, 5}

	s := Constructor(input)

	if !reflect.DeepEqual(s.nums, input) {
		panic("invalid Solution initilization")
	}

	if r := s.Reset(); !reflect.DeepEqual(r, input) {
		panic("invalid Solution.Reset() result")
	}

	if r := s.Shuffle(); !compareUnordered(r, input) {
		panic("invalid Solution.Shuffle() result")
	}
}

func compareUnordered(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}

	exists := make(map[int]int)

	for i := 0; i < len(first); i++ {
		v1, v2 := first[i], second[i]

		exists[v1]++
		exists[v2]--
	}

	for _, v := range exists {
		if v != 0 {
			return false
		}
	}

	return true
}
