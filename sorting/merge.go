package sorting

import (
	"fmt"

	"github.com/practice/algorithm/lib"
)

//MergeMain unit test for merge sort
func MergeMain() {
	input := lib.GenerateRandInts(10)
	fmt.Println(input)
	mrg := merge{}
	mrg.sort(input)
	fmt.Println("--------------------")
	fmt.Println(input)
}

type merge struct {
}

func (m merge) sort(input []int) {
	aux := make([]int, len(input))
	m.sorting(input, aux, 0, len(input)-1)
}

func (m merge) sorting(input []int, aux []int, low int, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	m.sorting(input, aux, low, mid)
	m.sorting(input, aux, mid+1, high)
	m.merging(input, aux, low, mid, high)
}

func (m merge) merging(input []int, aux []int, low int, mid int, high int) {
	//copy to aux array
	for index, val := range input {
		aux[index] = val
	}
	//merge to main array
	i := low
	j := mid + 1
	for k := low; k <= high; k++ {
		if i > mid {
			input[k] = aux[j]
			j++

		} else if j > high {
			input[k] = aux[i]
			j++

		} else if aux[j] < aux[i] {
			input[k] = aux[j]
			j++
		} else {
			input[k] = aux[i]
			i++
		}

	}
}
