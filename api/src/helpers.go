package src

import (
	"encoding/json"
	"net/http"
)

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
