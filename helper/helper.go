package helper

type response struct {
	Status  int
	Message string
	Data    interface{}
}

func ApiResponse(message string, status int, data interface{}) response {
	responseData := response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return responseData
}
