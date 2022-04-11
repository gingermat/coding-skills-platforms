package quick_union_weighted

type WeightedQU struct {
	ids  []int
	size []int // count objects in the tree rooted ids[i]
}

func NewWeightedQU(num int) *WeightedQU {
	ids := make([]int, num)
	sizes := make([]int, num)

	for i := range ids {
		ids[i] = i
		sizes[i] = 1
	}

	return &WeightedQU{ids: ids, size: sizes}
}

func (wu *WeightedQU) Union(p, q int) {
	i := wu.root(p)
	j := wu.root(q)

	if i == j {
		return
	}

	if wu.size[i] > wu.size[j] {
		wu.ids[i] = j
		wu.size[i] += wu.size[j]
	} else {
		wu.ids[j] = i
		wu.size[j] += wu.size[i]
	}
}

func (wu *WeightedQU) root(i int) int {
	for i != wu.ids[i] {
		i = wu.ids[i]
	}

	return i
}

func (wu *WeightedQU) IsConnected(p, q int) bool {
	return wu.root(p) == wu.root(q)
}
