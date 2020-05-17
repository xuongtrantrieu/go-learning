package exercises

import (
	"fmt"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func isMultipleOf(x int) func(num int) bool {
	var f = func(num int) bool {
		return num%x == 0
	}

	return f
}

func OddOrEven() {
	var err error

	isEven := isMultipleOf(2)
	isMultipleOf4 := isMultipleOf(4)

	var num int
	fmt.Printf("Enter a number: ")
	_, err = fmt.Scan(&num)
	utils.CheckErr(err)

	if isMultipleOf4(num) {
		fmt.Println("This is a multiple of 4")
	} else if isEven(num) {
		fmt.Println("This is an even")
	} else {
		fmt.Println("This is an odd")
	}

	var m int
	fmt.Printf("Input mother multiple of: ")
	_, err = fmt.Scan(&m)
	var customMultipleOf = isMultipleOf(m)

	var anotherNum int
	fmt.Printf("Insert to check if this number is multiple of %d: ", m)
	_, err = fmt.Scan(&anotherNum)

	if customMultipleOf(anotherNum) {
		fmt.Printf("%d is multiple of %d\n", anotherNum, m)
	} else {
		fmt.Printf("%d is not a multiple of %d\n", anotherNum, m)
	}
}
