package main

import "fmt"

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

func main() {
	q:=NewQueue(0)
	for i:=1;i<10;i++ {
		q.Push(i)
		fmt.Println(q.Pop())
	}
	fmt.Println(q.Pop())
}
