package exercises

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func WriteToAFile() {
	var err error

	fmt.Printf("Url: ")
	var url string
	_, err = fmt.Scan(&url)
	utils.CheckErr(err)

	fmt.Printf("Filepath: ")
	var filePath string
	_, err = fmt.Scan(&filePath)
	utils.CheckErr(err)

	response, err := http.Get(url)
	utils.CheckErr(err)

	defer func() {
		err := response.Body.Close()
		utils.CheckErr(err)
	}()

	rawBody, err := ioutil.ReadAll(response.Body)
	utils.CheckErr(err)

	err = ioutil.WriteFile(filePath, rawBody, 0644)
	utils.CheckErr(err)

	absFilePath, err := filepath.Abs(filePath)
	utils.CheckErr(err)

	fmt.Printf("File created: %s\n", absFilePath)
}
