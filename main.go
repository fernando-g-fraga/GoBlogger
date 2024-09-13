package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

type slugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct{}

type PostData struct {
	Title   string
	Content template.HTML
	Author  string
}

func (fsr FileReader) Read(slug string) (string, error) {
	f, err := os.Open(slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := io.ReadAll(f)

	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PostHandler(sl slugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postMarkdown, err := sl.Read(slug)

		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		var buf bytes.Buffer

		mdConverter := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("dracula"),
				),
			),
		)

		tpl, err := template.ParseFiles("post.gohtml")
		if err != nil {
			http.Error(w, "Error parsing the HTML files", http.StatusInternalServerError)
			return
		}

		err = mdConverter.Convert([]byte(postMarkdown), &buf)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, PostData{
			Title:   "Must-Have Items for Cat Owners",
			Content: template.HTML(buf.String()),
			Author:  "Fernando Fraga",
		})

	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReader{}))

	fmt.Println("Listening to port :3030")
	err := http.ListenAndServe(":3030", mux)

	if err != nil {
		log.Fatal(err)
	} else {

	}

}
