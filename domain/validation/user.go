package validation

import (
	"errors"
	"regexp"
)

var (
	ErrTooLongEmail = errors.New("email address is too long")
	ErrInvalidEmail = errors.New("email address is invalid")
)

func ValidateEmail(email string) error {
	if len(email) > 256 {
		return ErrTooLongEmail
	}

	const emailPattern string = `^[a-zA-Z0-9_+-.]+@[a-z0-9-.]+\.[a-z]+$`
	var emailRegexp *regexp.Regexp = regexp.MustCompile(emailPattern)
	if !emailRegexp.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

var (
	ErrTooShortPassword = errors.New("password is too short")
	ErrTooLongPassword  = errors.New("password is too long")
	ErrInvalidPassword  = errors.New("password is invalid")
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrTooShortPassword
	}

	if len(password) > 256 {
		return ErrTooLongPassword
	}

	const pattern string = `^[a-zA-Z0-9_+-.]+$`
	var passwordRegexp *regexp.Regexp = regexp.MustCompile(pattern)
	if !passwordRegexp.MatchString(password) {
		return ErrInvalidPassword
	}

	return nil
}
