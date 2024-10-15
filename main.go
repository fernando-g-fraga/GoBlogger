package main

import (
	"html/template"
	"io"
	"net/http"

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

	return c.Render(http.StatusOK, "contato.html", map[string]bool{
		"feito": true,
	})
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
	e.POST("/enviar_contato", send_contact)

	e.Logger.Fatal(e.Start(":8080"))

}
