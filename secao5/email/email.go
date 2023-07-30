package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func SendEmail(to []string, subject string, servidor string, erro string, horario string, templatePath string) {
	from := "roberto.m0010@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		panic("GMAIL_PASSWORD não foi configurada")
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles(templatePath)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  servidor,
		Error:   erro,
		Horario: horario,
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Printf("Erro ao enviar o email: %s", err)
		return
	}
	fmt.Println("Email enviado com sucesso!")

}
