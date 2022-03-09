package src

import (
	"encoding/json"
	"net/http"
)

func removeBook(s []Book, i int) []Book {
	if i != len(s)-1 {
		s[i] = s[len(s)-1]
	}
	return s[:len(s)-1]
}

func checkDuplicateBookId(s []Book, id int) bool {
	for _, book := range s {
		if book.Id == id {
			return true
		}
	}
	return false
}

func JSONResponse(w http.ResponseWriter, code int, desc string) {
	w.WriteHeader(code)
	message, ok := responseCodes[code]
	if !ok {
		message = "Undefined"
	}
	r := CustomResponse{
		Code:        code,
		Message:     message,
		Description: desc,
	}
	_ = json.NewEncoder(w).Encode(r)
}
