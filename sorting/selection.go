package sorting

import (
	"fmt"
	"strings"

	"github.com/practice/algorithm/lib"
)

//SelectionMain unit test
func SelectionMain() {
	fileName := "tiny.txt"
	//fileName := "words3.txt"
	output, err := lib.ReadAllString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	input := make([]string, 0, 1)
	for _, value := range output {
		rows := strings.Split(value, " ")
		for _, v := range rows {
			input = append(input, v)
		}
	}
	fmt.Println("Input Array:", input)
	fmt.Println("Size:", len(input))
	sel := &selection{}
	sel.sortStrings(input)
	fmt.Println("Sorted Array:", input)
}

type selection struct {
}

func (s selection) sortStrings(inputs []string) {
	for i := 0; i < len(inputs); i++ {
		minIndex := i
		for j := i + 1; j < len(inputs); j++ {
			if inputs[j] < inputs[minIndex] {
				minIndex = j
			}
		}
		temp := inputs[i]
		inputs[i] = inputs[minIndex]
		inputs[minIndex] = temp
	}
}
