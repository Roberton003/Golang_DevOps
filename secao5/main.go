package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func main() {
	from := "seumeail@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		panic("GMAIL_PASSWORD não foi configurada")
	}
	to := []string{
		"meaildedestino@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Alerta: Servidor down \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  "Google",
		Error:   "Erro ao acessar o servidor.",
		Horario: "24/07/2023 14:00",
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Printf("Erro ao enviar o email: %s", err)
		return
	}
	fmt.Println("Email enviado com sucesso!")

}
