package unionFind

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/practice/algorithm/lib"
)

//WeightedQuickUnionMain unit testing
func WeightedQuickUnionMain() {
	fileName := "tinyUF.txt"
	output, err := lib.ReadAllString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	qu := &weightQuickUnion{}
	for index, value := range output {
		if index == 0 {
			count, _ := strconv.Atoi(strings.TrimSpace(value))
			qu.init(count)
			continue
		}
		value = strings.TrimSpace(value)
		rows := strings.Split(value, " ")
		p, _ := strconv.Atoi(strings.TrimSpace(rows[0]))
		q, _ := strconv.Atoi(strings.TrimSpace(rows[1]))
		conn, err := qu.connected(p, q)
		if err != nil {
			fmt.Println(err)
		}
		if !conn {
			qu.union(p, q)
			fmt.Println(p, "and", q, "connected")
		}
	}
	fmt.Println("Connected Component:", qu.count)
	fmt.Println("Final array:", qu.parent)
	fmt.Println("Rank array:", qu.rank)
}

type weightQuickUnion struct {
	parent []int
	rank   []int //rank represent depth of each element
	count  int
}

func (uf *weightQuickUnion) init(n int) {
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	uf.count = n
}

func (uf weightQuickUnion) size() int {
	return uf.count
}

func (uf weightQuickUnion) find(p int) (int, error) {
	err := uf.validate(p)
	if err != nil {
		return -1, err
	}
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p, err
}

func (uf weightQuickUnion) validate(p int) (err error) {
	if p < 0 || p >= len(uf.parent) {
		msg := fmt.Sprintf("Illegal Argument Exception index %d is not in range between 0 and %d", p, len(uf.parent)-1)
		err = errors.New(msg)
	}
	return err
}

func (uf weightQuickUnion) connected(p, q int) (bool, error) {
	rootP, err := uf.find(p)
	if err != nil {
		return false, err
	}
	rootQ, err := uf.find(q)
	if err != nil {
		return false, err
	}
	return rootP == rootQ, err
}

func (uf *weightQuickUnion) union(p, q int) (err error) {
	rootP, err := uf.find(p)
	if err != nil {
		return err
	}
	rootQ, err := uf.find(q)
	if err != nil {
		return err
	}
	if rootP == rootQ {
		msg := fmt.Sprintf("%d and %d are already in the same group", p, q)
		err = errors.New(msg)
		return err
	}
	//make smaller root point to larger one
	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
		uf.rank[rootQ] += uf.rank[rootP]
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP] += uf.rank[rootQ]
	}
	uf.count--
	return err
}
