package models

type messageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) messageResponse {
	return messageResponse{message}
}
