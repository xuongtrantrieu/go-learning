package exercises

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func CowsAndBulls() {
	rand.Seed(time.Now().UnixNano())
	correctNumber := rand.Intn(10000-1000) + 1000

	correctChars := []byte(strconv.Itoa(correctNumber))

	playing := true

	containInList := func(l []byte, c byte) bool {
		for _, x := range l {
			if c == x {
				return true
			}
		}
		return false
	}

	var guess int
	for playing {
		fmt.Printf("Guess: ")
		_, err := fmt.Scan(&guess)
		utils.CheckErr(err)

		rawGuessChars := []byte(strconv.Itoa(guess))
		guessChars := make([]byte, 4)
		for i, c := range rawGuessChars {
			guessChars[i] = c
		}

		cowCount, bullCount := 0, 0
		for i, c := range correctChars {
			if containInList(guessChars, c) {
				if guessChars[i] == c {
					cowCount++
				} else {
					bullCount++
				}
			}
		}

		fmt.Printf("%d cow, %d bull\n", cowCount, bullCount)

		if cowCount == 4 {
			playing = false
		}
	}
}
