package exercises

import (
	"fmt"
	"time"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func CharInput() {
	var err error

	var name string
	fmt.Printf("Enter your name: ")
	_, err = fmt.Scan(&name)
	utils.CheckErr(err)

	var age int
	fmt.Printf("Enter your age: ")
	_, err = fmt.Scan(&age)
	utils.CheckErr(err)

	remainingYears := 100 - age
	theYear := remainingYears + time.Now().Year()

	message := fmt.Sprintf("Hello %v, you will turn 100 in: %d\n", name, theYear)
	fmt.Print(message)

	var printTimes int
	fmt.Printf("How many times do you want to reprint: ")
	_, err = fmt.Scan(&printTimes)

	for i := 0; i < printTimes; i++ {
		fmt.Printf("%d: %s", i, message)
	}
}
