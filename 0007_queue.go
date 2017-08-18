package main

import (
	"errors"
	"fmt"
)

func queueMain() {
	que := &queue{}
	que.enque("ajay")
	que.enque("sanjay")
	que.enque("priya")
	que.enque("raj")
	fmt.Println("size:", que.size())
	peekElement, _ := que.peek()
	fmt.Println("peek:", peekElement)
	que.dequeue()
	fmt.Println("After dequeue")
	fmt.Println("size:", que.size())
	peekElement, _ = que.peek()
	fmt.Println("peek:", peekElement)
	itr := que.iterator()
	for itr.hasNext() {
		item, _ := itr.next()
		fmt.Println(item)
	}
}

type queue struct {
	first *Node
	last  *Node
	count int
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}

func (q *queue) size() int {
	return q.count
}

func (q *queue) peek() (Item, error) {
	if q.isEmpty() {
		return nil, errors.New("No such element exception -queue underflow")
	}
	item := q.first.item
	return item, nil
}

func (q *queue) enque(item Item) {
	oldLast := q.last
	q.last = &Node{
		item: item,
	}
	if q.isEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.count++
}

func (q *queue) dequeue() (item Item, err error) {
	if q.isEmpty() {
		err = errors.New("No such element exception -queue underflow")
		return item, err
	}
	item = q.first.item
	q.first = q.first.next
	if q.isEmpty() {
		q.last = nil
	}
	return item, err
}

func (q *queue) iterator() Iterator {
	return &queueIterator{
		current: q.first,
	}
}

type queueIterator struct {
	current *Node
}

func (i *queueIterator) hasNext() bool {
	return i.current != nil
}

func (i *queueIterator) next() (item Item, err error) {
	if !i.hasNext() {
		err = errors.New("No such element exception")
		return item, err
	}
	item = i.current.item
	i.current = i.current.next
	return item, err
}

func (i *queueIterator) remove() (bool, error) {
	return false, errors.New("Unsupported operation")
}
