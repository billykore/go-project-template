package greet

import "gorm.io/gorm"

// Greet is greet entity.
type Greet struct {
	gorm.Model
	Name string
}

// EmailData represents the structure of an email message,
// containing a subject and body content.
type EmailData struct {
	// Subject is the title of the email.
	Subject string
	// Body is the main content of the email message.
	Body string
}
