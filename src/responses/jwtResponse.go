package responses

// Exception is an exception.
type Exception struct {
	Message string `json:"message"`
}

// Response is a response.
type Response struct {
	Data string `json:"data"`
}
