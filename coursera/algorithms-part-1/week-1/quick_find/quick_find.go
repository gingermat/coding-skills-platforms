package quick_find

type QF struct {
	ids []int
}

func NewQF(num int) *QF {
	ids := make([]int, num)

	for i := range ids {
		ids[i] = i
	}

	return &QF{ids: ids}
}

func (qf *QF) Union(p, q int) {
	pID := qf.ids[p]
	qID := qf.ids[q]

	for i := 0; i < len(qf.ids); i++ {
		if qf.ids[i] == pID {
			qf.ids[i] = qID
		}
	}
}

func (qf *QF) IsConnected(p, q int) bool {
	return qf.ids[p] == qf.ids[q]
}
