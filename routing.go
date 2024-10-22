package main

import (
	"github.com/fernandogfaga/GoBlogger/controller"
	"github.com/labstack/echo/v4"
)

/*
	e.GET("/", controller.HomeHandler)
	e.GET("/contato", contato)
	e.GET("/curriculo", curriculo)
	e.GET("/blog", blog)
	e.GET("/blog_admin", blog_admin)
	e.GET("/blog/:id", blog_page)
	e.POST("/enviar_contato", send_contact)
	e.POST("/enviar_blog", send_blog)
*/

//Handlers

func Gethome(e *echo.Echo) {
	e.GET("/", controller.HomeHandler)
}
func Getcontato(e *echo.Echo) {
	e.GET("/contato", controller.ContatoHandler)
}
func Getcurriculo(e *echo.Echo) {
	e.GET("/curriculo", controller.CurriculoHandler)
}
func Getblog(e *echo.Echo) {
	e.GET("/blog", controller.BlogHandler)
}
func GetblogAdmin(e *echo.Echo) {
	e.GET("/blog_admin", controller.BlogAdminHandler)
}
func GetBlogByID(e *echo.Echo) {
	e.GET("/blog_admin", controller.BlogIDHandler)
}
func POSTContato(e *echo.Echo) {
	e.POST("/enviar_contato", controller.Enviarcontact)
}
func POSTBlog(e *echo.Echo) {
	e.POST("/blog/enviar_blog/", controller.Enviarblog)
}
