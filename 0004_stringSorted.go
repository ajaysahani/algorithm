package main

import (
	"fmt"
	"strings"
)

func stringSortedMain() {
	str := "abcd"
	sorted := stringSorted{}.isStringSorted(str)
	fmt.Println(str, " is sorted:", sorted)
}

type stringSorted struct {
}

func (s stringSorted) isStringSorted(str string) bool {
	str = strings.ToLower(str)
	size := len(str)
	for i := 1; i < size-1; i++ {
		if str[i] < str[i-1] {
			return false
		}
	}
	return true
}
