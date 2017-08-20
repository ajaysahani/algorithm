package unionFind

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/practice/algorithm/lib"
)

//UfMain unit test
func UfMain() {
	fileName := "tinyUF.txt"
	output, err := lib.ReadAllString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	qu := &uf{}
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

type uf struct {
	parent []int
	rank   []int
	count  int
}

func (u *uf) init(n int) (err error) {
	if n < 0 {
		err = errors.New("Illegal argument exception")
		return err
	}
	u.parent = make([]int, n)
	u.rank = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
	}
	u.count = n
	return err
}

func (u uf) size() int {
	return u.count
}

func (u uf) find(p int) (int, error) {
	err := u.validate(p)
	if err != nil {
		return -1, err
	}
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]] //path compression by halving
		p = u.parent[p]
	}
	return p, err
}

func (u uf) validate(p int) (err error) {
	if p < 0 || p >= len(u.parent) {
		msg := fmt.Sprintf("Illegal Argument Exception index %d is not in range between 0 and %d", p, len(u.parent)-1)
		err = errors.New(msg)
	}
	return err
}

func (u uf) connected(p, q int) (bool, error) {
	rootP, err := u.find(p)
	if err != nil {
		return false, err
	}
	rootQ, err := u.find(q)
	if err != nil {
		return false, err
	}
	return rootP == rootQ, err
}

func (u *uf) union(p, q int) (err error) {
	rootP, err := u.find(p)
	if err != nil {
		return err
	}
	rootQ, err := u.find(q)
	if err != nil {
		return err
	}
	if rootP == rootQ {
		msg := fmt.Sprintf("%d and %d are already in the same group", p, q)
		err = errors.New(msg)
		return err
	}

	if u.rank[rootP] < u.rank[rootQ] {
		u.parent[rootP] = rootQ
	} else if u.rank[rootP] > u.rank[rootQ] {
		u.parent[rootQ] = rootP
	} else {
		u.parent[rootQ] = rootP
		u.rank[rootP]++
	}
	u.count--
	return err
}
