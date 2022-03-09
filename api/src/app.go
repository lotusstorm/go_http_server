package src

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/book", getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/book", addBook).Methods(http.MethodPost)
	router.HandleFunc("/book/{book_id:[0-9]+}", getBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{book_id:[0-9]+}", updateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{book_id:[0-9]+}", deleteBook).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:5000", handlers.LoggingHandler(os.Stdout, router)))
}
