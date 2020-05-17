package exercises

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func parseSize(s string) (int, int) {
	var width, height int
	var err error

	for i, x := range strings.Split(s, "x") {
		if i == 0 {
			width, err = strconv.Atoi(x)
			utils.CheckErr(err)
		}

		if i == 1 {
			height, err = strconv.Atoi(x)
			utils.CheckErr(err)
			break
		}
	}

	return width, height*2
}

func DrawAGameBoard() {
	evenLine := " ---"
	oddLine := "|   "

	var boardSize string
	fmt.Printf("Enter board size: ")
	_, err := fmt.Scan(&boardSize)
	utils.CheckErr(err)

	width, height := parseSize(boardSize)
	utils.CheckErr(err)

	for i := 0; i < height+1; i++ {
		var printed string

		if i % 2 == 0 {
			printed = evenLine
		} else {
			printed = oddLine
		}

		for j := 0; j < width; j++ {
			fmt.Printf(printed)
		}

		if i % 2 != 0 {
			fmt.Println(string(oddLine[0]))
		} else {
			fmt.Println()
		}
	}
}
