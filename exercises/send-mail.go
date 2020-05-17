package exercises

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/xuongtrantrieu/go-learning/utils"
)

func SendMail() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("to: ")
	rawTos, err := reader.ReadString('\n')
	utils.CheckErr(err)

	tos := make([]string, 0)
	for _, t := range strings.Split(strings.Trim(rawTos, "\n"), ";") {
		tos = append(tos, strings.Trim(t, " "))
	}

	auth := smtp.PlainAuth("", "jenkins@geoguard.com", "D34FG4rrttyyR4", "smtp.gmail.com")
	subject := "Subject: Test sending email via golang\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	message := []byte(subject + mime + "Sent via <b>golang</b>")

	err = smtp.SendMail("smtp.gmail.com:587", auth, "jenkins@gmail.com", tos, message)
	utils.CheckErr(err)
}
