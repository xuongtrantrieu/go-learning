package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func ReverseWordOrder() {
	fmt.Printf("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	utils.CheckErr(err)
	text = strings.Trim(text, "\n")

	words := strings.Split(text, " ")
	reversedWords := make([]string, 0)
	for i := len(words) - 1; i > -1; i-- {
		reversedWords = append(reversedWords, words[i])
	}

	reversedText := strings.Join(reversedWords, " ")
	fmt.Printf("Reversed sentence: %s\n", reversedText)
}
