package lib

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	rootPath = "/data/"
)

func getFilePath(fileName string) string {
	dir, _ := os.Getwd()
	path := dir + rootPath + fileName
	return path
}

//ReadAllString read strings from file
func ReadAllString(fileName string) (output []string, err error) {
	path := getFilePath(fileName)
	file, err := os.Open(path)
	if err != nil {
		return output, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	output = make([]string, 0, 1)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output, err
}

//ReadAllInt reads all int from file
func ReadAllInt(fileName string) (output []int, err error) {
	out, err := ReadAllString(fileName)
	if err != nil {
		return output, err
	}
	size := len(out)
	output = make([]int, size)
	for index, value := range out {
		value = strings.TrimSpace(value)
		mValue, _ := strconv.Atoi(value)
		output[index] = mValue
	}
	return output, err
}

//GenerateRandInts generate random ints
func GenerateRandInts(count int) []int {
	input := make([]int, count)
	for i := 0; i < count; i++ {
		input[i] = rand.Intn(99)
	}
	return input
}
