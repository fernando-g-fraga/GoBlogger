package util

import (
	"fmt"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(name, email, message string) error {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading the env variables")
	}

	var b string

	b = fmt.Sprintf("Você recebeu um novo contato! \n a pessoa %s, e-mail%s deixou a seguinte mensagem \n %s", name, email, message)

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

func MDToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
