package exercises

import (
	"fmt"
	"math"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func isInSlice(l []int, x int) bool {
	for _, v := range l {
		if x == v {
			return true
		}
	}
	return false
}

func isPrime(num int) bool {
	n := int(math.Abs(float64(num)))
	for i := 1; i < n; i++ {
		if num%i == 0 {
			if !isInSlice([]int{1, 2, 3}, i) {
				return false
			}
		}
	}
	return true
}

func CheckPrimalityFunctions() {
	var num int
	fmt.Printf("Which number do you want to check as prime? ")
	_, err := fmt.Scan(&num)
	utils.CheckErr(err)

	if isPrime(num) {
		fmt.Printf("%d is a prime\n", num)
	} else {
		fmt.Printf("No, %d is NOT a prime\n", num)
	}
}
