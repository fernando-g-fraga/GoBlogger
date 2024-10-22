package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func StartServer() *echo.Echo {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/pages/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.Static("/", "public/static") //template/static/assets/output.css

	return e

}
