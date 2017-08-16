package main

import "fmt"

func tpMain() {
	t := &mtest{
		count: 5,
	}
	fmt.Println(t.get())
	t.decrement()
	fmt.Println(t.get())
	t.decrement()
	fmt.Println(t.get())
	t.decrement()
	fmt.Println(t.get())
}

type mtest struct {
	count int
}

func (t mtest) decrement() {
	t.count--
}

func (t mtest) get() int {
	t.count--
	val := t.count
	return val
}
