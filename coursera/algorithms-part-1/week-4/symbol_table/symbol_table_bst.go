package symbol_table

type SymbolTableBST struct {
	root *BinarySearchTree
}

func (s *SymbolTableBST) Get(k Key) Value {
	return s.root.Get(k)
}

func (s *SymbolTableBST) Put(k Key, v Value) {
	s.root.Put(k, v)
}

func (s *SymbolTableBST) Delete(k Key) {
	s.root.Delete(k)
}

func (s *SymbolTableBST) Contains(k Key) bool {
	return s.Get(k) != nil
}

func (s *SymbolTableBST) Size() int {
	return s.root.Size()
}

func (s *SymbolTableBST) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SymbolTableBST) Keys() []Key {
	return s.root.Keys()
}

func NewSymbolTableBST() *SymbolTableBST {
	return &SymbolTableBST{
		root: NewBinarySearchTree(),
	}
}
