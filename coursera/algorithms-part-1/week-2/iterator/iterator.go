package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Iterable interface {
	Iterator() *Iterator
}
