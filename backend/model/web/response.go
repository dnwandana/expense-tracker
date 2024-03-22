package web

// ResponseData is a struct that represents the response data format
type ResponseData struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

// ResponseMessage is a struct that represents the response message format
type ResponseMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
