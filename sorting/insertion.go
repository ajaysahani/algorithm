package sorting

import (
	"fmt"

	"github.com/practice/algorithm/lib"
)

//InsertionMain unit test
func InsertionMain() {
	input := lib.GenerateRandInts(10)
	fmt.Println(input)
	ins := insertion{}
	ins.sort(input)
	fmt.Println("--------------------")
	fmt.Println(input)
}

type insertion struct {
}

func (ins insertion) sort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := i; j > 0 && input[j] < input[j-1]; j-- {
			temp := input[j]
			input[j] = input[j-1]
			input[j-1] = temp
		}
	}
}
