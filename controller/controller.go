package controller

import (
	"net/http"

	"github.com/fernandogfaga/GoBlogger/model"
	util "github.com/fernandogfaga/GoBlogger/util"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "home.html", map[string]string{"Title": "Home"})
}
func ContatoHandler(c echo.Context) error {
	// api_key := os.Getenv("Google_Key")
	data := map[string]string{
		"title": "Contato",
		// "API_KEY": api_key,
	}
	return c.Render(http.StatusOK, "contato.html", data)
}

func CurriculoHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "curriculo.html", nil)
}
func BlogHandler(c echo.Context) error {
	return nil
}
func BlogAdminHandler(c echo.Context) error {
	return nil
}
func BlogIDHandler(c echo.Context) error {
	return nil
}
func Enviarcontact(c echo.Context) error {
	name := c.FormValue("nome")
	email := c.FormValue("email")
	mensagem := c.FormValue("mensagem")

	util.SendMail(name, mensagem, email)

	return c.Redirect(http.StatusFound, "/contato")
}

func Enviarblog(c echo.Context) error {
	var post model.Blog_post

	// post.ID = rand.Intn(999)
	post.Title = c.FormValue("titulo")
	post.Post = c.FormValue("post")
	post.Tags = append(post.Tags, c.FormValue("tags"))

	cod, msg := model.CreateBlogPost(post)
	return c.Render(cod, "blog_admin.html", msg)
}
