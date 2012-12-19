package main

import "fmt"
import "errors"

type LinkedList struct {
	first *link
	last *link
}

func (r *LinkedList) AddEnd(val int) {
	p:=link{r.last,nil,val}
	r.last.next=&p
	r.last=&p
}

func (r *LinkedList) AddBeg(val int) {
	p:=link{nil,r.first,val}
	r.first.prev=&p
	r.first=&p
}

func (r *LinkedList) Iterator() Iterator {
	return Iterator{&link{nil,r.first,0}}
}

func NewLinkedList(firstVal int) LinkedList {
	p:=link{nil,nil,firstVal}
	return LinkedList{&p,&p}
}

type link struct {
	prev *link
	next *link
	value int
}

type Iterator struct {
	pointer *link
}

func (it *Iterator) Next() (int, error) {
	if it.pointer==nil||it.pointer.next==nil {
		return 0,errors.New("Iterator: no next value")
	}
	it.pointer=it.pointer.next
	return it.pointer.value, nil
}

func (it *Iterator) Prev() (int,error) {
	if it.pointer==nil||it.pointer.prev==nil {
		return 0,errors.New("Iterator: no next value")
	}
	it.pointer=it.pointer.prev
	return it.pointer.value,nil
}
