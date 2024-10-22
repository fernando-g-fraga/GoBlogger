package model

import (
	"log"
	"os"
)

type Blog_post struct {
	ID    int
	Title string
	Post  string
	Tags  []string
}

func CreateBlogPost(p Blog_post) (int, string) {
	path := "public/blog/"

	fs, err := os.OpenFile(path+p.Title+".md", os.O_RDWR, 0777)

	if err != nil {
		log.Println(err)
		return 500, "Erro ao criar o post do blog."
	}

	_, err = fs.Write([]byte(p.Post))

	if err != nil {
		log.Println(err)
		return 500, "Erro ao salvar o arquivo do blog."
	}

	return 200, "Post Criado com sucesso!"
}
