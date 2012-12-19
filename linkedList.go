package main

import "fmt"

type LinkedList struct {
	First *Link
	Last *Link
}

func (r *LinkedList) addEnd(val int) {
	p:=Link{r.Last,nil,val}
	r.Last.Next=&p
	r.Last=&p
}

func (r *LinkedList) addBeg(val int) {
	fmt.Println("Adding ",val," to ", r.First)
	p:=Link{nil,r.First,val}
	r.First.Prev=&p
	r.First=&p
}

func (r *LinkedList) iterator() Iterator {
	return Iterator{&Link{nil,r.First,0}}
}

func newReadyQueue(firstVal int) LinkedList {
	p:=Link{nil,nil,firstVal}
	return LinkedList{&p,&p}
}

type Link struct {
	Prev *Link
	Next *Link
	value int
}

type Iterator struct {
	I *Link
}

func (it *Iterator) next() *Link {
	if it.I==nil||it.I.Next==nil {
		return nil
	}
	it.I=it.I.Next
	return it.I
}

func (it *Iterator) prev() *Link {
	if it.I==nil||it.I.Prev==nil {
		return nil
	}
	it.I=it.I.Prev
	return it.I
}

func main() {
	count:=10
	rq:=newReadyQueue(20)
	for i:=count-1;i>=0;i-- {
		rq.addBeg(i)
		rq.addEnd(i)
	}
	it:=rq.iterator()
	for i:=0; i<1; i++ {
		for p:=it.next();p!=nil;p=it.next() {
			fmt.Println(*p)
		}
	}
}
