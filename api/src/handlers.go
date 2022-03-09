package src

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, _ := strconv.Atoi(vars["book_id"])
	for _, book := range booksDB {
		if book.Id == bookId {
			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}
	JSONResponse(w, http.StatusNotFound, "")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(booksDB)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var b Book
	err := decoder.Decode(&b)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "")
		return
	}
	if checkDuplicateBookId(booksDB, b.Id) {
		JSONResponse(w, http.StatusConflict, "")
		return
	}
	booksDB = append(booksDB, b)
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(b)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, _ := strconv.Atoi(vars["book_id"])
	decoder := json.NewDecoder(r.Body)
	var b Book
	err := decoder.Decode(&b)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "")
		return
	}
	for i, book := range booksDB {
		if book.Id == bookId {
			booksDB[i] = b
			_ = json.NewEncoder(w).Encode(b)
			return
		}
	}
	JSONResponse(w, http.StatusNotFound, "")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, _ := strconv.Atoi(vars["book_id"])
	for i, book := range booksDB {
		if book.Id == bookId {
			booksDB = removeBook(booksDB, i)
			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}
	JSONResponse(w, http.StatusNotFound, "")
}
