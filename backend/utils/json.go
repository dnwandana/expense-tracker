package utils

import (
	"encoding/json"
	"net/http"
)

// ReadJsonRequest is a function that decodes the JSON body of the HTTP request into the provided empty interface as the result.
func ReadJsonrequest(request *http.Request, result any) {
	err := json.NewDecoder(request.Body).Decode(result)
	PanicIfError(err)
}

// WriteJsonResponse is a function that writes the provided response as JSON to the HTTP response writer.
func WriteJsonResponse(writer http.ResponseWriter, response any, statusCode ...int) {
	writer.Header().Set("Content-Type", "application/json")
	status := http.StatusOK
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(response)
	PanicIfError(err)
}
