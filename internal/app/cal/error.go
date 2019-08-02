package cal

import (
	"errors"
	"net/http"
)

// Errors
var (
	ErrEmailRequired      = errors.New("auth: email required")
	ErrEmailInvalid       = errors.New("auth: email invalid")
	ErrEmailNotVerify     = errors.New("auth: email not verify")
	ErrEmailNotAvailable  = errors.New("auth: email not available")
	ErrPasswordRequired   = errors.New("auth: password required")
	ErrPasswordInvalid    = errors.New("auth: password invalid")
	ErrInvalidCredentials = errors.New("auth: invalid credentials")
	ErrTokenRequired      = errors.New("auth: token required")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrEmailRequired, ErrEmailInvalid, ErrEmailNotVerify, ErrEmailNotAvailable:
		return http.StatusBadRequest
	case ErrPasswordRequired, ErrPasswordInvalid:
		return http.StatusBadRequest
	case ErrInvalidCredentials:
		return http.StatusUnauthorized
	case ErrTokenRequired:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrEmailRequired:
		return "email required"
	case ErrEmailInvalid:
		return "email invalid"
	case ErrEmailNotAvailable:
		return "email not available"
	case ErrEmailNotVerify:
		return "please verify email before sign in"
	case ErrPasswordRequired:
		return "password required"
	case ErrPasswordInvalid:
		return "password length must be 8 - 200 characters"
	case ErrInvalidCredentials:
		return "email or password invalid"
	case ErrTokenRequired:
		return "token required"
	default:
		return "Internal Server Error"
	}
}
