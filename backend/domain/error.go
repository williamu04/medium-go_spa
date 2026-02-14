package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrEmailAlreadyUsed  = errors.New("email already used")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrInvalidPassword   = errors.New("invalid password (minimum 8 characters)")
	ErrCredentialInvalid = errors.New("invalid credentials")
	ErrInternalError     = errors.New("internal server error")
	ErrMissingUserID     = errors.New("missing user ID")
	ErrMissingFields     = errors.New("missing required fields")
	ErrDatabaseOpen      = errors.New("failed to open database")
	ErrDatabasePing      = errors.New("failed to ping database")
	ErrInvalidTopic      = errors.New("invalid topic (minimum 1)")
	ErrTopicNotFound     = errors.New("topic not found")
)
