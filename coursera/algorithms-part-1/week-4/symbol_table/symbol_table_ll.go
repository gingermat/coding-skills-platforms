package symbol_table

type Node struct {
	key   Key
	value Value
	next  *Node
}

type SymbolTableLL struct {
	first *Node
	size  int
}

func (s *SymbolTableLL) Get(k Key) Value {
	for node := s.first; node != nil; node = node.next {
		if node.key == k {
			return node.value
		}
	}
	return nil
}

func (s *SymbolTableLL) Put(k Key, v Value) {
	for node := s.first; node != nil; node = node.next {
		if node.key == k {
			node.value = v
			return
		}
	}

	s.size++
	s.first = &Node{k, v, s.first}
}

func (s *SymbolTableLL) Delete(k Key) {
	var prev *Node

	for curr := s.first; curr != nil; prev, curr = curr, curr.next {
		if curr.key != k {
			continue
		}

		if prev == nil {
			s.first = curr.next
		} else {
			prev.next = curr.next
		}

		s.size--
	}
}

func (s *SymbolTableLL) Contains(k Key) bool {
	return s.Get(k) != nil
}

func (s *SymbolTableLL) Size() int {
	return s.size
}

func (s *SymbolTableLL) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SymbolTableLL) Keys() []Key {
	var keys []Key

	for node := s.first; node != nil; node = node.next {
		keys = append(keys, node.key)
	}

	return keys
}

func NewSymbolTableLL() *SymbolTableLL {
	return &SymbolTableLL{}
}
