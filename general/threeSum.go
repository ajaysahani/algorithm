package general

import (
	"fmt"

	"github.com/practice/algorithm/lib"
)

//ThreeSumMain unit test
func ThreeSumMain() {
	fileName := "1Kints.txt"
	output, _ := lib.ReadAllInt(fileName)
	tsum := &threeSum{}
	tsum.printAll(output)
	count := tsum.count(output)
	fmt.Println("Count: ", count)
}

type threeSum struct {
}

func (t *threeSum) printAll(input []int) {
	size := len(input)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if input[i]+input[j]+input[k] == 0 {
					fmt.Println(input[i], " ", input[j], " ", input[k])
				}
			}
		}
	}
}

func (t *threeSum) count(input []int) int {
	size := len(input)
	count := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if input[i]+input[j]+input[k] == 0 {
					count++
				}
			}
		}
	}
	return count
}
