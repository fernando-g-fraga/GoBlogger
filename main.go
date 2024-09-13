package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/post/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		fmt.Fprintf(w, "Post: %s", slug)
	})

	fmt.Println("Listening to port :3030")
	err := http.ListenAndServe(":3030", mux)

	if err != nil {
		log.Fatal(err)
	} else {

	}

}
