package user

import (
	"context"

	"github.com/billykore/go-service-tmpl/internal/pkg/constant"
)

// User represents a system user with identification and contact information.
type User struct {
	ID    int
	CIF   string
	Name  string
	Email string
	Phone string
}

// FromContext gets user data from ctx context.
func FromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(constant.UserContextKey).(*User)
	return user, ok
}
