package controller

type outputMessage struct {
	Message string `json:"message"`
}

var invalidJsonFormatError = &outputMessage{
	Message: "Invalid JSON format",
}

var successMessage = &outputMessage{
	Message: "Success",
}

var internalServerError = &outputMessage{
	Message: "Internal Server Error",
}

var bookNotFound = &outputMessage{
	Message: "Book not found",
}
