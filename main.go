package main

import (
	"fmt"

	"github.com/practice/algorithm/lib"
)

func main() {
	testLib()
	//tpMain()
	//gcdMain()
	//binarySearchMain()
	//pailndromeMain()
	//stringSortedMain()
	//bagMain()
	//stackMain()

}

func testLib() {
	fileName := "1Kints.txt"
	output, err := lib.ReadAllInt(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(output))
}
