package quick_union

type QU struct {
	ids []int
}

func NewQU(num int) *QU {
	ids := make([]int, num)

	for i := range ids {
		ids[i] = i
	}

	return &QU{ids: ids}
}

func (qu *QU) Union(p, q int) {
	j := qu.root(p)
	i := qu.root(q)

	qu.ids[i] = j
}

func (qu *QU) root(i int) int {
	for i != qu.ids[i] {
		i = qu.ids[i]
	}

	return i
}

func (qu *QU) IsConnected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}
