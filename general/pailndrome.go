package general

import (
	"fmt"
	"strings"
)

//PailndromeMain unit test
func PailndromeMain() {
	checkString := "Nitin"
	output := pailndrome{}.isPailndrome(checkString)
	fmt.Println(checkString, " is palindrome:", output)
}

type pailndrome struct {
}

func (p pailndrome) isPailndrome(str string) bool {
	size := len(str)
	str = strings.ToLower(str)
	for i := 0; i < size/2; i++ {
		if str[i] != str[size-1-i] {
			return false
		}
	}
	return true
}
