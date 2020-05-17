package exercises

import (
	"fmt"

	"github.com/anaskhan96/soup"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func DecodeAWebPage() {
	var err error

	url := "http://nhattruyen.com/"
	//url := "https://xkcd.com/"
	resp, err := soup.Get(url)
	utils.CheckErr(err)

	doc := soup.HTMLParse(resp)

	for i, titleRoot := range doc.FindAll("a", "class", "jtip") {
		fmt.Printf("#%d - %s\n", i+1, titleRoot.Text())
	}
}
