package main

import "fmt"

func gcdMain() {
	fmt.Println("Hello GCD")
	p := 2
	q := 0
	op := gcd(p, q)
	fmt.Printf("GCD for %d and %d= %d", p, q, op)
	fmt.Println("")
}

func gcd(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}
