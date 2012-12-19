package main

type Queue struct {
	list *LinkedList
}

func NewQueue(val int) Queue {
	ll:=NewLinkedList(val)
	return Queue{&ll}
}

func (q *Queue) Push(val int) {
	q.list.AddEnd(val)
}

func (q *Queue) Pop() (int,error) {
	return q.list.RemoveBeg()
}
