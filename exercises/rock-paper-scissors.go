package exercises

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/xuongtrantrieu/go-learning/utils"
)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"
)

func getComputerMove() string {
	rand.Seed(time.Now().UnixNano())

	var move string
	switch moveNumber := rand.Int() % 3; moveNumber {
	case 0:
		move = ROCK
	case 1:
		move = PAPER
	default:
		move = SCISSORS
	}

	return move
}

func getUserMove() string {
	var err error

	var move string
	fmt.Printf("What is your move(rock/paper/scissors)? ")
	_, err = fmt.Scan(&move)
	move = strings.ToLower(move)
	utils.CheckErr(err)

	switch move {
	case "rock":
		fallthrough
	case "paper":
		fallthrough
	case "scissors":
		err = nil
	default:
		err = errors.New("invalid move")
	}
	utils.CheckErr(err)

	return move
}

func compareMoves(a, b string) int {
	var result int

	switch a {
	case ROCK:
		switch b {
		case ROCK:
			result = -1
		case PAPER:
			result = 1
		default:
			result = 0
		}
	case PAPER:
		switch b {
		case ROCK:
			result = 0
		case PAPER:
			result = -1
		default:
			result = 1
		}
	default:
		switch b {
		case ROCK:
			result = 1
		case PAPER:
			result = 0
		default:
			result = -1
		}
	}

	return result
}

func RockPaperScissors() {
	playing := true
	var confirmation string

	for playing {
		yourMove := getUserMove()
		computerMove := getComputerMove()

		fmt.Printf("Your move: %s\n", yourMove)
		fmt.Printf("Computer's move: %s\n", computerMove)

		result := compareMoves(yourMove, computerMove)
		if result == -1 {
			fmt.Println("Draw!")
		} else if result == 0 {
			fmt.Println("You win!")
		} else {
			fmt.Println("You lose")
		}

		fmt.Printf("\nContinue playing?(y/n) ")
		_, err := fmt.Scan(&confirmation)
		utils.CheckErr(err)

		switch confirmation {
		case "y":
			playing = true
		case "n":
			playing = false
		default:
			utils.CheckErr(errors.New("invalid confirmation"))
		}
	}
	fmt.Println("Goodbye!")

}
