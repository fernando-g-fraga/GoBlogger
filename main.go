package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type slugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct{}

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
		markdown, err := sl.Read(slug)

		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, markdown)
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
