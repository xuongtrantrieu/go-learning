package exercises

import (
	"fmt"

	"github.com/xuongtrantrieu/go-learning/utils"
)

var A = [...]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

func ListLessThanTen() {
	fmt.Printf("Original list: %v\n", A)

	var lessThan5List = make([]int, 0)
	for _, x := range A {
		if !(x < 5) {
			continue
		}
		lessThan5List = append(lessThan5List, x)
	}

	fmt.Printf("List less than 5: %v\n", lessThan5List)

	var limit int
	fmt.Printf("Define your limit: ")
	_, err := fmt.Scan(&limit)
	utils.CheckErr(err)

	customLessThanList := make([]int, 0)
	for _, x := range A {
		if x < limit {
			customLessThanList = append(customLessThanList, x)
		}
	}
	fmt.Printf("List less than %d: %v\n", limit, customLessThanList)
}
