package helper

import (
	"Go-UserManagement/model/web"
	"encoding/json"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	response := web.WebResponse{
		Code:   code,
		Status: "OK",
		Data:   data,
	}
	WriteToResponseBody(w, response)
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	response := web.WebResponse{
		Code:   code,
		Status: "ERROR",
		Data:   message,
	}
	WriteToResponseBody(w, response)
}

func WriteToResponseBody(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}