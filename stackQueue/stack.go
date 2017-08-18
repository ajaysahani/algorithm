package stackQueue

import (
	"errors"
	"fmt"
)

//StackMain unit testing
func StackMain() {
	s := &stack{}
	s.push("ajay")
	s.push("sanjay")
	s.push("raj")
	s.push("priya")
	fmt.Println("size:", s.size())
	item, _ := s.peek()
	fmt.Println("peek:", item)
	s.pop()
	fmt.Println("After Pop-----")
	fmt.Println("size:", s.size())
	item, _ = s.peek()
	fmt.Println("peek:", item)
	itr := s.getIterator()
	for itr.hasNext() {
		fmt.Println(itr.next())
	}
}

type stack struct {
	first *Node
	count int
}

func (s *stack) isEmpty() bool {
	return s.first == nil
}

func (s *stack) size() int {
	return s.count
}

func (s *stack) push(item Item) {
	node := &Node{
		item: item,
		next: s.first,
	}
	s.first = node
	s.count++
}

func (s *stack) pop() (Item, error) {
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	s.count--
	return item, nil
}

func (s *stack) peek() (Item, error) {
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}
	item := s.first.item
	return item, nil
}

func (s *stack) getIterator() Iterator {
	return &stackIterator{
		current: s.first,
	}
}

type stackIterator struct {
	current *Node
}

func (i *stackIterator) hasNext() bool {
	return i.current != nil
}

func (i *stackIterator) next() (Item, error) {
	if !i.hasNext() {
		return nil, errors.New("No such element exception")
	}
	item := i.current.item
	i.current = i.current.next
	return item, nil
}

func (i *stackIterator) remove() (bool, error) {
	return false, errors.New("Unsupported operation")
}
