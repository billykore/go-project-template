package email

import (
	"github.com/billykore/go-service-tmpl/domain/greet"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewGreetEmail, wire.Bind(new(*GreetEmail), new(greet.Email)),
)
