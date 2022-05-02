package queue

type UnorderedMaxPQ struct {
	pq []int
	n  int
}

func (q *UnorderedMaxPQ) IsEmpty() bool {
	return q.n == 0
}

func (q *UnorderedMaxPQ) Insert(x int) {
	n := q.n
	q.pq[n], q.n = x, n+1
}

func (q *UnorderedMaxPQ) DelMax() int {
	var max int // index of max element

	for i := 1; i < q.n; i++ {
		if q.pq[max] < q.pq[i] {
			max = i
		}
	}

	maxValue := q.pq[max]
	n := q.n - 1

	q.n = n
	q.pq[max], q.pq[n] = q.pq[n], q.pq[max]

	return maxValue
}

func NewUnorderedMaxPQ(capacity int) *UnorderedMaxPQ {
	return &UnorderedMaxPQ{
		pq: make([]int, capacity),
	}
}
