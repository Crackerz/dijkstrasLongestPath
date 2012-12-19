package main

import "fmt"
import "errors"

type LinkedList struct {
	first *link
	last *link
}

func (r *LinkedList) AddEnd(val int) {
	p:=link{r.last,nil,val}
	if r.first==nil&&r.last==nil {
		r.first=&p
		r.last=&p
		return
	}
	r.last.next=&p
	r.last=&p
}

//This function removes a node from the end of the list
//returning its value upon successfully removing it
func (r *LinkedList) RemoveEnd() int, error {
	if r.last==nil {
		return 0, errors.New("LinkedList: Nothing To Remove")
	} else if r.last.prev==nil {
		val:=r.last.value
		r.first=nil
		r.last=nil
		return val,nil
	}
	val:=last.value
	r.last=r.last.prev
	r.last.next=nil
	return val,nil
}

func (r *LinkedList) AddBeg(val int) {
	p:=link{nil,r.first,val}
	if r.first==nil&&r.last==nil {
		r.first=&p
		r.last=&p
		return
	}
	r.first.prev=&p
	r.first=&p
}

//This function removes a node from the end of the list
//returning its value upon successfully removing it
func (r *LinkedList) RemoveBeg() int, error {
	if r.first==nil {
		return 0,errors.New("LinkedList: Nothing To Remove")
	} else if r.first.next==nil {
		val:=first.value
		r.first=nil
		r.last=nil
		return val,nil
	}
	val:=first.value
	r.first=r.first.next
	r.first.prev=nil
	return val,nil
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
