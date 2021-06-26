package main

type Queue struct {
	items chan int
}

type Item struct {
	value int
}

func (q *Queue) Enqueue(item int) {
	q.items <- item
}

func (q *Queue) Dequeue() int {
	return <-q.items
}

func main() {
	q := Queue{items: make(chan int, 3)}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}
