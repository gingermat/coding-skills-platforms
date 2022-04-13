package stack

const (
	MINIMAL_STACK_CAPACITY = 2 // minimal items for initialize stack

	STACK_GROW_FACTOR      float64 = 2.0  // increase array on STACK_GROW_FACTOR times
	STACK_SHRINK_FACTOR    float64 = 0.5  // shrink array on STACK_SHRINK_FACTOR times
	STACK_OCCUPANCY_FACTOR float64 = 1.0  // increase array when filled >= STACK_OCCUPANCY_FACTOR
	STACK_EMPTYNESS_FACTOR float64 = 0.25 // shrink array where filled < STACK_EMPTYNESS_FACTOR
)

type StackOnArray struct {
	idx   int
	items []interface{} // developed on slice, but array cannot defined with non-constant value
}

func (st *StackOnArray) Push(item interface{}) {
	if st.fillRatio() >= STACK_OCCUPANCY_FACTOR {
		newSize := int(float64(len(st.items)) * STACK_GROW_FACTOR)
		st.resize(newSize)
	}

	st.idx++
	st.items[st.idx] = item
}

func (st *StackOnArray) Pop() interface{} {
	item := st.items[st.idx]
	st.idx--

	if st.fillRatio() <= STACK_EMPTYNESS_FACTOR {
		newSize := int(float64(len(st.items)) * STACK_SHRINK_FACTOR)
		st.resize(newSize)
	}

	return item
}

func (st *StackOnArray) IsEmpty() bool {
	return st.idx == -1
}

func (st *StackOnArray) resize(size int) {
	newSized := make([]interface{}, size)
	copy(newSized, st.items)

	st.items = newSized
}

func (st *StackOnArray) fillRatio() float64 {
	return float64(st.idx+1) / float64(len(st.items))
}

func NewStackOnArray() *StackOnArray {
	return &StackOnArray{
		items: make([]interface{}, MINIMAL_STACK_CAPACITY),
		idx:   -1,
	}
}
