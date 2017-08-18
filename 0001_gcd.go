package main

import "fmt"

func gcdMain() {
	fmt.Println("Hello GCD")
	p := 2
	q := 0
	op := gcd{}.greatedCommonDivisor(p, q)
	fmt.Printf("GCD for %d and %d= %d", p, q, op)
	fmt.Println("")
}

type gcd struct {
}

func (c gcd) greatedCommonDivisor(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return c.greatedCommonDivisor(q, r)
}
