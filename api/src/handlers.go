package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func (app *App) checkBookExists(field string, value string) bool {
// 	var count int64
// 	err := app.DB.Model(&Book{}).Select("id").Where(fmt.Sprintf("%s = ?", field), value).Count(&count).Error
// 	if err != nil {
// 		return false
// 	}
// 	if count > 0 {
// 		return true
// 	}
// 	return false
// }

// func (app *App) getBook(w http.ResponseWriter, r *http.Request) {
// 	var book Book
// 	vars := mux.Vars(r)
// 	bookId, _ := strconv.Atoi(vars["book_id"])
// 	result := app.DB.First(&book, bookId)
// 	if result.Error == nil {
// 		_ = json.NewEncoder(w).Encode(book)
// 		return
// 	}
// 	JSONResponse(w, http.StatusNotFound, "")
// }

// func (app *App) getAllBooks(w http.ResponseWriter, r *http.Request) {
// 	var books []Book
// 	_ = app.DB.Order("id asc").Find(&books)
// 	_ = json.NewEncoder(w).Encode(books)
// }

// func (app *App) addBook(w http.ResponseWriter, r *http.Request) {
// 	var book Book
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&book)
// 	if err != nil {
// 		JSONResponse(w, http.StatusBadRequest, "")
// 		return
// 	}
// 	if app.checkBookExists("title", book.Title) {
// 		JSONResponse(w, http.StatusConflict, "")
// 		return
// 	}
// 	result := app.DB.Create(&book)
// 	if result.Error == nil {
// 		w.WriteHeader(201)
// 		_ = json.NewEncoder(w).Encode(book)
// 	}
// }

// func (app *App) updateBook(w http.ResponseWriter, r *http.Request) {
// 	var book Book
// 	var newBook Book
// 	vars := mux.Vars(r)
// 	bookId, _ := strconv.Atoi(vars["book_id"])
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&newBook)
// 	if err != nil {
// 		JSONResponse(w, http.StatusBadRequest, "")
// 		return
// 	}
// 	if app.checkBookExists("title", newBook.Title) {
// 		JSONResponse(w, http.StatusConflict, "")
// 		return
// 	}
// 	result := app.DB.First(&book, bookId)
// 	if result.Error == nil {
// 		newBook.Id = bookId
// 		app.DB.Save(&newBook)
// 		_ = json.NewEncoder(w).Encode(newBook)
// 		return
// 	}
// 	JSONResponse(w, http.StatusNotFound, "")
// }

// func (app *App) deleteBook(w http.ResponseWriter, r *http.Request) {
// 	var book Book
// 	vars := mux.Vars(r)
// 	bookId, _ := strconv.Atoi(vars["book_id"])
// 	result := app.DB.First(&book, bookId)
// 	if result.Error == nil {
// 		app.DB.Delete(&book)
// 		return
// 	}
// 	JSONResponse(w, http.StatusNotFound, "")
// }

// **************Quiz**************

func (app *App) addQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz Quiz
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&quiz)
	fmt.Println(quiz)

	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "")
		return
	}
	// if app.checkBookExists("title", book.Title) {
	// 	JSONResponse(w, http.StatusConflict, "")
	// 	return
	// }
	result := app.DB.Create(&quiz)
	if result.Error == nil {
		w.WriteHeader(201)
		_ = json.NewEncoder(w).Encode(quiz)
	}

	// app.DB.Create(&quiz)

	// var quizs []Quiz
	// app.DB.Find(&quizs)
	// fmt.Println(quizs)
}

func (app *App) getAllQuizs(w http.ResponseWriter, r *http.Request) {
	var quizs []Quiz
	_ = app.DB.Order("id asc").Find(&quizs)
	_ = json.NewEncoder(w).Encode(quizs)
}
