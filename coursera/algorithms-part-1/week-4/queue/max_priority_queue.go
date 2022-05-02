package queue

type MaxPriorityQueue struct {
	pq []int
	n  int
}

func (q *MaxPriorityQueue) IsEmpty() bool {
	return q.n == 0
}

func (q *MaxPriorityQueue) Insert(x int) {
	n := q.n + 1

	q.pq[n], q.n = x, n
	q.swim(q.n)
}

func (q *MaxPriorityQueue) DelMax() int {
	max := q.pq[1]

	q.exch(1, q.n)
	q.n--

	q.sink(1)
	q.pq[q.n+1] = 0

	return max
}

func (q *MaxPriorityQueue) swim(k int) {
	for ; k > 1 && q.less(k/2, k); k /= 2 {
		q.exch(k, k/2)
	}
}

func (q *MaxPriorityQueue) sink(k int) {
	for 2*k <= q.n {
		j := 2 * k

		if j < q.n && q.less(j, j+1) {
			j++
		}

		if !q.less(k, j) {
			break
		}

		q.exch(k, j)
		k = j
	}
}

func (q *MaxPriorityQueue) exch(i int, j int) {
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
}

func (q *MaxPriorityQueue) less(i int, j int) bool {
	return q.pq[i] < q.pq[j]
}

func NewMaxPriorityQueue(capacity int) *MaxPriorityQueue {
	return &MaxPriorityQueue{
		pq: make([]int, capacity),
	}
}
