package utilities

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func SendCodeMail(receiver, code string) error {

	//TODO: Change mail template html to inner style not from css
	//TODO: Logout after a certain time
	//TODO: AuthSession cookie

	from := GetSettings().Mail
	password := GetSettings().Password

	to := []string{receiver}

	smtpHost := GetSettings().ServerSMTP
	smtpPort := GetSettings().PortSMTP

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles("./static/codemail.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Docs - Codice di verifica \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Code string
	}{Code: code})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		panic(err)
	}
	return nil
}
