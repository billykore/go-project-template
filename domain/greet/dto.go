package greet

type Response struct {
	Name string `json:"name"`
}

type HelloRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type HelloResponse struct {
	Message string `json:"message"`
}
