package exercises

import (
	"fmt"
	"math/rand"
	"time"
)

func removeDuplicateFromList(l []int) []int {
	var contains = func(l []int, x int) bool {
		for _, v := range l {
			if x == v {
				return true
			}
		}
		return false
	}

	result := make([]int, 0)
	for _, v := range l {
		if contains(result, v) {
			continue
		}

		result = append(result, v)
	}

	return result
}

func ListRemoveDuplicates() {
	rand.Seed(time.Now().UnixNano())
	l := make([]int, 0)
	for i := 0; i < 20; i++ {
		l = append(l, rand.Intn(20))
	}

	fmt.Printf("Initilized list: %v\n", l)
	fmt.Printf("Duplicates removed list: %v\n", removeDuplicateFromList(l))
}
