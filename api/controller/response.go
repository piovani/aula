package controller

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseMessageError struct {
	Error string `json:"error"`
}

func NewResponseMessage(msg string) ResponseMessage {
	return ResponseMessage{
		Message: msg,
	}
}

func NewResponseMessageError(msg string) ResponseMessageError {
	return ResponseMessageError{
		Error: msg,
	}
}
