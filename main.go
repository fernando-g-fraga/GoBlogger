package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fernandogfaga/GoBlogger/util"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func curriculo(c echo.Context) error {
	return c.Render(http.StatusOK, "curriculo.html", nil)
}

func blog_admin(c echo.Context) error {
	return c.Render(http.StatusOK, "blog_admin.html", nil)
}

func blog_page(c echo.Context) error {
	id := c.Param("id")
	log.Println(id)
	basepath := "template/static/blog/"
	md, err := os.ReadFile(basepath + id + ".md")

	if err != nil {
		log.Println("Page not found.")
	}

	html := util.MDToHTML(md)

	return c.HTML(http.StatusOK, string(html))
}

func blog(c echo.Context) error {
	var all_files []string

	files, err := os.ReadDir("template/static/blog/")

	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		log.Println(file.Name())
		all_files = append(all_files, file.Name())
	}

	log.Println(all_files)

	return c.Render(http.StatusOK, "blog.html", map[string][]string{
		"s": all_files,
	})

}

func main() {
	//load env
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading env variables")
	}
	//start Server
	e := StartServer()
	//Routing
	Gethome(e)
	Getcurriculo(e)
	Getcontato(e)
	Getblog(e)
	GetblogAdmin(e)
	GetBlogByID(e)
	POSTBlog(e)
	POSTContato(e)
	//Run Server
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port must be set")
	}
	e.Logger.Fatal(e.Start(":" + port))

}
