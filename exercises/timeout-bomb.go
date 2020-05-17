package exercises

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func getAnswer(answer chan string) {
	reader := bufio.NewReader(os.Stdin)

	rawAnswer, err := reader.ReadString('\n')
	utils.CheckErr(err)

	answer <- strings.Trim(rawAnswer, "\n")
}

func setupBomb(boom chan bool) {
	bomb := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-bomb.C:
			fmt.Println("You are too late!")
			boom <- true
			return
		}
	}
}

func question(answer chan string) string {
	bomb := time.NewTicker(5 * time.Second)

	for {
		select {
		case a := <-answer:
			bomb.Stop()
			return a
		case <-bomb.C:
			fmt.Println("You're late!")
			return "..."
		}
	}
}

func TimeoutBomb() {
	rawFileContent, err := ioutil.ReadFile("/Users/xuong/projects/go-learning/assets/timeout-bomb.txt")
	utils.CheckErr(err)
	fileContent := string(rawFileContent)

	questions := strings.Split(
		strings.Trim(fileContent, "\n"),
		"\n",
	)

	qaa := make(map[string]string)
	for i, q := range questions {
		fmt.Printf("Question %d: %s ", i+1, q)

		answerChan := make(chan string)

		go getAnswer(answerChan)

		answer := question(answerChan)
		qaa[q] = answer
	}

	fmt.Println(qaa)
}
