package email

import (
	"context"

	"github.com/billykore/go-service-tmpl/domain/greet"
)

type GreetEmail struct{}

func NewGreetEmail() *GreetEmail {
	return &GreetEmail{}
}

func (ge *GreetEmail) Send(ctx context.Context, data greet.EmailData) error {
	return nil
}
