package exercises

import (
	"fmt"
	"strings"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func isPalindrome(s string) bool {
	charOfS := strings.Split(s, "")

	reversedChars := make([]string, 0)
	for _, c := range charOfS {
		reversedChars = append([]string{c}, reversedChars...)
	}

	reversedS := strings.Join(reversedChars, "")

	if strings.Compare(s, reversedS) == 0 {
		return true
	}

	return false
}

func StringLists() {
	var text string
	fmt.Printf("Check palindrome: ")
	_, err := fmt.Scan(&text)
	utils.CheckErr(err)

	if isPalindrome(text) {
		fmt.Printf("%s is a palindrome\n", text)
	} else {
		fmt.Printf("No, %s is not a palindrome\n", text)
	}
}
