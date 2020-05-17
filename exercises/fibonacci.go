package exercises

import (
	"fmt"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func fibonacci(n int) []int {
	a, b := 1, 1
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		result = append(result, a)
		a, b = b, a+b
	}

	return result
}

func Fibonacci() {
	var num int
	fmt.Printf("How many Fibonacci? ")
	_, err := fmt.Scan(&num)
	utils.CheckErr(err)

	fmt.Printf("Your fibonacci: %v\n", fibonacci(num))
}
