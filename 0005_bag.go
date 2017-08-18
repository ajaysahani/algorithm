package main

import (
	"errors"
	"fmt"
)

func bagMain() {
	bg := &bag{}
	bg.add("ajay")
	bg.add("Rachana")
	bg.add("Trisha")
	fmt.Println("size:", bg.size())
	it := bg.iterator()
	for it.hasNext() {
		fmt.Println(it.next())
	}
}

type bag struct {
	first *Node
	count int
}

func (b *bag) isEmpty() bool {
	return b.first == nil
}

func (b *bag) size() int {
	return b.count
}

func (b *bag) add(item Item) {
	oldFirst := b.first
	node := &Node{
		item: item,
		next: oldFirst,
	}
	b.first = node
	b.count++
}

func (b *bag) iterator() Iterator {
	return &bagIterator{current: b.first}
}

type bagIterator struct {
	current *Node
}

func (b *bagIterator) hasNext() bool {
	return b.current != nil
}

func (b *bagIterator) next() (Item, error) {
	if b.hasNext() == false {
		return nil, errors.New("No such element exception")
	}
	item := b.current.item
	b.current = b.current.next
	return item, nil
}

func (b *bagIterator) remove() (bool, error) {
	return false, errors.New("Operation not supported")
}
