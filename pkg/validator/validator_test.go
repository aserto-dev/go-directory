package validator_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aserto-dev/go-directory/pkg/validator"
	"github.com/stretchr/testify/assert"
)

type tc struct {
	value    string
	expected error
}

// message: "must be all lowercase, start with a letter, can contain letters, digits, dots, underscores, and dashes, and must end with a letter or digit".
// expression: "this.matches('^[a-z][a-z0-9\\\\._-]{1,62}[a-z0-9]$')".
var typeIdentifierTests = []tc{
	{
		value:    "", // no type specified
		expected: nil,
	},
	{
		value:    "aaa", // min length
		expected: nil,
	},
	{
		value:    "aaa4567890123456678901234566789012345667890123456678901234567890", // max length
		expected: nil,
	},
}

func TestTypeIdentifier(t *testing.T) {
	for i, tc := range typeIdentifierTests {
		err := validator.TypeIdentifier(tc.value)
		if tc.expected == nil {
			assert.NoError(t, err, "test %d", i)
		} else {
			assert.Error(t, tc.expected, err, "test %d", i)
		}
	}
}

// message: "cannot contain any spaces or other whitespace characters".
// expression: "this.matches('^[\\\\S]+$')".
var instanceIdentifierTests = []tc{
	{
		value:    "aaa", // min length
		expected: nil,
	},
	{
		value:    fmt.Sprintf("a23456%s", strings.Repeat("1234567890", 25)), // max length (256)
		expected: nil,
	},
	{
		value:    fmt.Sprintf("a234567%s", strings.Repeat("1234567890", 25)), // max length exceeded (256+)
		expected: validator.ErrLengthInstanceIdentifier,
	},
}

func TestInstanceIdentifier(t *testing.T) {
	for _, tc := range instanceIdentifierTests {
		err := validator.InstanceIdentifier(tc.value)
		if tc.expected == nil {
			assert.NoError(t, err)
		} else {
			assert.Error(t, tc.expected, err)
		}
	}
}

var displayNameTests = []tc{
	{
		value:    "aaa", // min length
		expected: nil,
	},
	{
		value:    strings.Repeat("1234567890", 10), // max length (100)
		expected: nil,
	},
	{
		value:    "a" + strings.Repeat("1234567890", 10), // exceeds max length (100+)
		expected: validator.ErrLengthDisplayName,
	},
	{
		value:    fmt.Sprintf("%s \t \n ", "Hello World"),
		expected: nil,
	},
}

func TestDisplayName(t *testing.T) {
	for _, tc := range displayNameTests {
		err := validator.DisplayName(tc.value)
		if tc.expected == nil {
			assert.NoError(t, err)
		} else {
			assert.Error(t, tc.expected, err)
		}
	}
}

// digits only (Etag is a stringified int64).
// default: "0".
// max_len: 20 characters/digits.
var etagTests = []tc{
	{
		value:    "0", // min length & default value
		expected: nil,
	},
	{
		value:    strings.Repeat("1234567890", 2), // max length (20)
		expected: nil,
	},
	{
		value:    strings.Repeat("1234567890", 2) + "1", // larger then max length (20+)
		expected: validator.ErrLengthEtag,
	},
	{
		value:    "abc", // no digits
		expected: validator.ErrInvalidEtag,
	},
	{
		value:    "123 456 678", // digits with spaces
		expected: validator.ErrInvalidEtag,
	},
}

func TestEtag(t *testing.T) {
	for _, tc := range etagTests {
		err := validator.DisplayName(tc.value)
		if tc.expected == nil {
			assert.NoError(t, err)
		} else {
			assert.Error(t, tc.expected, err)
		}
	}
}
