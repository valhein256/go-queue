//+build disable

package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

type Item struct {
	value int
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	for i, e := range q.items {
		println(i, e)
	}
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}
