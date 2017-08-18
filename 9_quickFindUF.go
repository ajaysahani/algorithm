package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/practice/algorithm/lib"
)

func quickFindUFMain() {
	fileName := "tinyUF.txt"
	output, err := lib.ReadAllString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	qf := &quickFindUF{}
	for index, value := range output {
		if index == 0 {
			count, _ := strconv.Atoi(strings.TrimSpace(value))
			qf.init(count)
			continue
		}
		value = strings.TrimSpace(value)
		rows := strings.Split(value, " ")
		p, _ := strconv.Atoi(strings.TrimSpace(rows[0]))
		q, _ := strconv.Atoi(strings.TrimSpace(rows[1]))
		if !qf.connected(p, q) {
			qf.union(p, q)
			fmt.Println(p, "and", q, "connected")
		}
	}
	fmt.Println("Connected Component:", qf.count)
	fmt.Println("Final array:", qf.id)
}

type quickFindUF struct {
	id    []int
	count int
}

func (uf *quickFindUF) init(size int) {
	uf.id = make([]int, size)
	uf.count = size
	for i := 0; i < size; i++ {
		uf.id[i] = i
	}
}

func (uf *quickFindUF) size() int {
	return uf.count
}

func (uf *quickFindUF) find(p int) (int, error) {
	err := uf.validate(p)
	if err != nil {
		return -1, err
	}
	return uf.id[p], err
}

func (uf *quickFindUF) validate(p int) (err error) {
	if p < 0 || p >= uf.count {
		msg := fmt.Sprintf("Illegal Argument Exception index %d is not in range between 0 and %d ", p, (uf.count - 1))
		err = errors.New(msg)
	}
	return err
}

func (uf *quickFindUF) connected(p, q int) bool {
	err := uf.validate(p)
	if err != nil {
		return false
	}
	err = uf.validate(q)
	if err != nil {
		return false
	}
	return uf.id[p] == uf.id[q]
}

func (uf *quickFindUF) union(p, q int) (err error) {
	err = uf.validate(p)
	if err != nil {
		return err
	}
	err = uf.validate(q)
	if err != nil {
		return err
	}
	pID := uf.id[p]
	qID := uf.id[q]

	if pID == qID {
		msg := fmt.Sprintf("%d and %d are already in same group %d", p, q, uf.id[p])
		err = errors.New(msg)
		return err
	}
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
	return err
}
