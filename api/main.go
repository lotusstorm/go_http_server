package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World</h1>")
		fmt.Println("Server is started...")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
