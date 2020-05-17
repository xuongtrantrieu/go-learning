package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func ListEnds() {
	var err error

	in := bufio.NewReader(os.Stdin)

	var text string
	fmt.Printf("Input a list of int, seperated by ',': ")
	text, err = in.ReadString('\n')
	utils.CheckErr(err)

	nums := make([]int, 0)
	for _, t := range strings.Split(text, ",") {
		num, err := strconv.Atoi(
			strings.Trim(
				strings.Trim(t, " "),
				"\n",
			),
		)
		utils.CheckErr(err)

		nums = append(nums, num)
	}

	result := []int{nums[0], nums[len(nums)-1]}

	fmt.Printf("Result is: %v\n", result)
}
