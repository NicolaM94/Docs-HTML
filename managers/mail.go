package managers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

// Function to generate a random code.
//
// Private.
func Codegen() string {
	var out string
	var alfabeth string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var random = rand.New(rand.NewSource(time.Now().Unix()))
	for n := 0; n < 6; n++ {
		out += string(alfabeth[random.Intn(len(alfabeth))])
	}
	return out
}

// Function to send the code for authorization
func SendCodeMail(receiver string, code string) error {

	var settings Settings = Settings{}.Populate()
	var sendto []string = []string{receiver}
	var auth = smtp.PlainAuth("Organization Nicola", settings.ServerMailAddress, settings.ServerMailPass, settings.ServerSMTPHost)

	t, _ := template.ParseFiles("./static/codemail.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: DocShelf - Code \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Code string
	}{
		Code: code,
	})

	err := smtp.SendMail(settings.ServerSMTPHost+":"+settings.ServerSMTPPort, auth, settings.ServerMailAddress, sendto, body.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	log.Default().Println("! Code mail sent")
	return nil
}
