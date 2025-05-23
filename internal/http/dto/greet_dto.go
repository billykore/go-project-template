package dto

type HelloRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func NewHelloResponse(message string) *HelloResponse {
	return &HelloResponse{
		Message: message,
	}
}
