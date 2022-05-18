package hash_table

const NumOfChains = 97

type Node struct {
	key   Key
	value Value
	next  *Node
}

type SeparateChainingHashST struct {
	st []Node
}

func (h SeparateChainingHashST) Put(k Key, v Value) {
	i := h.hash(k)

	for x := &h.st[i]; x.key != ""; x = x.next {
		if k == x.key {
			x.value = v
			return
		}
	}

	h.st[i] = Node{k, v, &h.st[i]}
}

func (h SeparateChainingHashST) Get(k Key) Value {
	i := h.hash(k)

	for x := &h.st[i]; x.key != ""; x = x.next {
		if k == x.key {
			return x.value
		}
	}

	return nil
}

func (st *SeparateChainingHashST) hash(k Key) int {
	var s uint32

	for _, r := range k {
		s = 31*s + uint32(r)
	}

	s &= 0x7fffffff
	s %= NumOfChains

	return int(s)
}

func NewSeparateChainingHashST() *SeparateChainingHashST {
	return &SeparateChainingHashST{
		st: make([]Node, NumOfChains),
	}
}
