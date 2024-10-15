package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/fernandogfaga/GoBlogger/util"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func home(c echo.Context) error {

	data := map[string]string{
		"title": "Home"}
	return c.Render(http.StatusOK, "home.html", data)
}

func contato(c echo.Context) error {
	data := map[string]string{
		"title": "Contato",
	}
	return c.Render(http.StatusOK, "contato.html", data)
}

func send_contact(c echo.Context) error {
	name := c.FormValue("nome")
	email := c.FormValue("email")
	mensagem := c.FormValue("mensagem")

	mailnotification("mail", name, mensagem, email)

	return c.Render(http.StatusOK, "contato.html", map[string]bool{
		"feito": true,
	})
}

func mailnotification(form string, name, message, email string) {
	switch {
	case form == "mail":
		bigString := fmt.Sprintf("Você recebeu um novo contato! \n a pessoa %s, e-mail%s deixou a seguinte mensagem \n %s", name, email, message)
		err := util.SendMail(bigString)

		if err != nil {
			log.Println("Error sending the email.")
		}
	}
}

func curriculo(c echo.Context) error {
	return c.Render(http.StatusOK, "curriculo.html", "")
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("template/pages/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/", "template/static") //template/static/assets/output.css
	e.GET("/", home)
	e.GET("/contato", contato)
	e.GET("/curriculo", curriculo)
	e.POST("/enviar_contato", send_contact)

	e.Logger.Fatal(e.Start(":8080"))

}
