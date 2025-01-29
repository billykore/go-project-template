package dto

// Response is a simple struct containing a name, used for basic data representation.
type Response struct {
	Name string `json:"name"`
}

// HelloRequest is the input for the SayHello method, with validation rules for the name.
type HelloRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

// HelloResponse is the output for the SayHello method, containing a greeting message.
type HelloResponse struct {
	Message string `json:"message"`
}
