package exercises

import (
	"fmt"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func Divisors() {
	var num int
	fmt.Printf("Divisors of which number do you want to get? ")
	_, err := fmt.Scan(&num)
	utils.CheckErr(err)

	var divisors = []int{}
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	fmt.Printf("Divisors of %d is %v\n", num, divisors)
}
