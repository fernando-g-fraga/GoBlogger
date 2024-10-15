package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(b string) error {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading the env variables")
	}

	user_mail := os.Getenv("user_mail")
	password_mail := os.Getenv("password_mail")

	m := gomail.NewMessage()
	m.SetHeader("From", "contato@fernandofraga.com.br")
	m.SetHeader("To", "fernandogrfraga@gmail.com")
	m.SetHeader("Subject", "Novo contato!")
	m.SetBody("text/html", b)

	d := gomail.NewDialer("smtp.hostinger.com", 465, user_mail, password_mail)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending the email", err)
		return err
	}
	return nil
}
