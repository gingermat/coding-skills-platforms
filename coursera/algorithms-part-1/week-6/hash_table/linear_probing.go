package hash_table

import "fmt"

const NumOfCells = 301

type LinearProbingHashST struct {
	keys   []Key
	values []Value
}

func (h LinearProbingHashST) Put(k Key, v Value) {
	var i int

	for i = h.hash(k); h.keys[i] != ""; i = (i + 1) % NumOfCells {
		fmt.Println("------", i)
		if h.keys[i] == k {
			break
		}
	}

	h.keys[i] = k
	h.values[i] = v
}

func (h LinearProbingHashST) Get(k Key) Value {
	for i := h.hash(k); h.keys[i] != ""; i = (i + 1) % NumOfCells {
		if h.keys[i] == k {
			return h.values[i]
		}
	}

	return nil
}

func (st *LinearProbingHashST) hash(k Key) int {
	var s uint32

	for _, r := range k {
		s = 31*s + uint32(r)
	}

	s &= 0x7fffffff
	s %= NumOfCells

	return int(s)
}

func NewLinearProbingHashST() *LinearProbingHashST {
	return &LinearProbingHashST{
		keys:   make([]Key, NumOfCells),
		values: make([]Value, NumOfCells),
	}
}
