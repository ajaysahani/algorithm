package unionFind

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/practice/algorithm/lib"
)

func ufMain() {
	fileName := "tinyUF.txt"
	output, err := lib.ReadAllString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	for index, value := range output {
		if index == 0 {
			continue
		}
		value = strings.TrimSpace(value)
		rows := strings.Split(value, " ")
		val1, _ := strconv.Atoi(strings.TrimSpace(rows[0]))
		val2, _ := strconv.Atoi(strings.TrimSpace(rows[1]))
		fmt.Println(val1, ":", val2)
	}
}

//uf union find
type uf struct {
}
