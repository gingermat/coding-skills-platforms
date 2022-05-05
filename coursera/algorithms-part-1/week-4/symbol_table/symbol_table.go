package symbol_table

type Key string
type Value interface{}

type SymbolTable interface {
	Get(Key) Value
	Put(Key, Value)
	Delete(Key)
	Contains(Key) bool
	Size() int
	IsEmpty() bool
	Keys() []Key
}
