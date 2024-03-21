package utils

import (
	"encoding/json"
	"net/http"

	"github.com/dnwandana/expense-tracker/model/web"
)

// ReadJsonRequest is a function that decodes the JSON body of the HTTP request into the provided empty interface as the result.
func ReadJsonrequest(request *http.Request, result any) {
	err := json.NewDecoder(request.Body).Decode(request)
	PanicIfError(err)
}

// WriteJsonResponse is a function that sets the content type of the HTTP response to "application/json".
// Then, it encodes the web.Response object into a JSON format and writes it into the HTTP response.
func WriteJsonResponse(writer http.ResponseWriter, response web.Response) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	PanicIfError(err)
}
