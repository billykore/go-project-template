package entity

import (
	"context"

	"github.com/billykore/go-service-tmpl/pkg/utils/constant"
)

const (
	DefaultName = "John Doe"
	ForbiddenID = 13
)

// User represents a system user with identification and contact information.
type User struct {
	ID    int
	CIF   string
	Name  string
	Email string
	Phone string
}

// UserFromContext gets user data from ctx context.
func UserFromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(constant.UserContextKey).(*User)
	return user, ok
}
