package exercises

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func GuessingGameOne() {
	var err error
	playing := true

	for playing {
		fmt.Println("Welcome to Guessing Game One")
		fmt.Println("We have generated a random number from 1 -> 20")

		rand.Seed(time.Now().UnixNano())
		correctAnswer := rand.Int()%20 + 1
		var answer int

		isCorrect := answer == correctAnswer
		count := 0

		for !isCorrect {
			count++
			fmt.Printf("Guess what the number is? ")
			_, err = fmt.Scan(&answer)
			utils.CheckErr(err)

			if answer == correctAnswer {
				fmt.Printf("Congratulation! You win! It took you %v time(s)\n", count)
				isCorrect = true
			} else if answer > correctAnswer {
				fmt.Println("Smaller baby!")
			} else {
				fmt.Println("Bigger baby!")
			}
		}

		var playAgain string
		fmt.Printf("Play again?(y/n) ")
		_, err = fmt.Scan(&playAgain)
		utils.CheckErr(err)

		if playAgain == "n" {
			playing = false
		}
		fmt.Println()
	}

	fmt.Println("Goodbye.")
}
