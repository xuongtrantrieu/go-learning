package exercises

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func LogMultiDes() {
	var err error

	fmt.Printf("Enter filepath to log to: ")
	var logFilePath string
	_, err = fmt.Scan(&logFilePath)
	utils.CheckErr(err)

	finalLogFilePath, err := filepath.Abs(logFilePath)
	utils.CheckErr(err)

	folderPath := filepath.Dir(finalLogFilePath)
	_ = os.Mkdir(folderPath, 0755)

	logFile, err := os.OpenFile(finalLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func() {
		err := logFile.Close()
		utils.CheckErr(err)
	}()
	utils.CheckErr(err)
	streamLogger := log.New(os.Stdout, "", log.LstdFlags)
	fileLogger := log.New(logFile, "", log.LstdFlags|log.LUTC|log.Llongfile)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter something to log: ")
		somethingToLog, err := reader.ReadString('\n')
		utils.CheckErr(err)

		if somethingToLog == "\n" {
			break
		}

		logMessage := strings.Trim(somethingToLog, "\n")

		streamLogger.Print(logMessage)
		fileLogger.Print(logMessage)
	}
}
