package exercises

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateRandomList(n int) []int {
	rand.Seed(time.Now().UnixNano())

	k := 20
	l := make([]int, 0)
	for i := 0; i < n; i++ {
		l = append(l, rand.Int()%k)
	}

	return l
}

func hasInList(l []int, x int) bool {
	for _, v := range l {
		if v == x {
			return true
		}
	}
	return false
}

func hasInCommon(a, b []int) []int {
	commonList := make([]int, 0)
	for _, x := range a {
		if !hasInList(b, x) {
			continue
		}

		if hasInList(commonList, x) {
			continue
		}

		commonList = append(commonList, x)
	}

	for _, x := range b {
		if !hasInList(a, x) {
			continue
		}

		if hasInList(commonList, x) {
			continue
		}

		commonList = append(commonList, x)
	}

	sort.Ints(commonList)

	return commonList
}

func ListOverLap() {
	listA := generateRandomList(5)
	listB := generateRandomList(5)
	commonList := hasInCommon(listA, listB)

	fmt.Printf("List A: %v\nList B: %v\n", listA, listB)
	fmt.Printf("Common list: %v\n", commonList)
}
