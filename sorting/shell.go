package sorting

import (
	"fmt"

	"github.com/practice/algorithm/lib"
)

//ShellMain unit test
func ShellMain() {
	input := lib.GenerateRandInts(10)
	fmt.Println(input)
	shl := shell{}
	shl.sort(input)
	fmt.Println("--------------------")
	fmt.Println(input)
}

type shell struct {
}

func (sh shell) sort(input []int) {
	h := 1
	for h < len(input)/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len(input); i++ {
			for j := i; j >= h && input[j] < input[j-h]; j -= h {
				temp := input[j]
				input[j] = input[j-h]
				input[j-h] = temp
			}
		}
		h = h / 3
	}
}
