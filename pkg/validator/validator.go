package validator

import (
	"errors"
	"regexp"
)

// TypeIdentifier:
// default: N/A
// min length: 3
// max length: 64
// format: ^[a-zA-Z][a-zA-Z0-9._-]{1,62}[a-zA-Z0-9]$
// desc: must start with a letter, can contain mixed case letters, digits, dots, underscores, and dashes, and must end with a letter or digit.
const maxTypeIdentifierLength int = 64

var (
	ErrInvalidTypeIdentifier = errors.New("must start with a letter, can contain mixed case letters, digits, dots, underscores, and dashes, and must end with a letter or digit")
	ErrLengthTypeIdentifier  = errors.New("max length of 64 characters exceeded")
	typeIdentifierMatch      = regexp.MustCompile(`(?m)^[a-z][a-z0-9._-]{1,62}[a-z0-9]$`)
)

func TypeIdentifier(ident string) error {
	if ident == "" {
		return nil
	}
	if len(ident) > maxTypeIdentifierLength {
		return ErrLengthTypeIdentifier
	}
	if !typeIdentifierMatch.MatchString(ident) {
		return ErrInvalidTypeIdentifier
	}
	return nil
}

// InstanceIdentifier:
// default: N/A
// min length: 1
// max length: 256
// format: ^[S]{1,256}$
// desc: cannot contain any spaces or other whitespace characters.
const maxInstanceIdentifierLength int = 256

var (
	ErrInvalidInstanceIdentifier = errors.New("cannot contain any spaces or other whitespace characters")
	ErrLengthInstanceIdentifier  = errors.New("max length of 256 characters exceeded")
	instanceIdentifierMatch      = regexp.MustCompile(`(?m)^\S{1,256}$`)
)

func InstanceIdentifier(ident string) error {
	if ident == "" {
		return nil
	}
	if len(ident) > maxInstanceIdentifierLength {
		return ErrLengthInstanceIdentifier
	}
	if !instanceIdentifierMatch.MatchString(ident) {
		return ErrInvalidInstanceIdentifier
	}
	return nil
}

// DisplayName:
// default: ""
// min length: 0
// max length: 100
// format: printable characters.
// desc: must not contain angled brackets, ampersands, or double quotes
const maxDisplayNameLength int = 100

var (
	ErrInvalidDisplayName = errors.New("display name can only contain printable characters")
	ErrLengthDisplayName  = errors.New("max length of 100 characters exceeded")
	displayNameMatch      = regexp.MustCompile(`(?m)^[[:print:]]{0,100}$`)
)

func DisplayName(dn string) error {
	if dn == "" {
		return nil
	}
	if len(dn) > maxDisplayNameLength {
		return ErrLengthDisplayName
	}
	if !displayNameMatch.MatchString(dn) {
		return ErrInvalidDisplayName
	}
	return nil
}

// Etag:
// default: "0"
// min length: 1
// max length: 20
// format: digits only.
const maxEtagLength int = 20

var (
	ErrInvalidEtag = errors.New("can only contain digits")
	ErrLengthEtag  = errors.New("max length of 20 characters exceeded")
	etagMatch      = regexp.MustCompile(`(?m)^\d{1,20}$`)
)

func Etag(etag string) error {
	if etag == "" {
		return nil
	}
	if len(etag) > maxEtagLength {
		return ErrLengthEtag
	}
	if !etagMatch.MatchString(etag) {
		return ErrInvalidEtag
	}
	return nil
}
