package validate_test

import (
	"errors"
	"testing"

	dsc3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	"github.com/bufbuild/protovalidate-go"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonValidation(t *testing.T) {
	msg := &dsc3.Object{
		Type: "user-helper",
		Id:   "euang@acmecorp.com",
	}

	v, err := protovalidate.New()
	require.NoError(t, err)
	if err = v.Validate(msg); err != nil {
		t.Logf("validation failed: %s", err)
	} else {
		t.Log("validation succeeded")
	}
}

func TestIdValidation(t *testing.T) {
	v, err := protovalidate.New()
	require.NoError(t, err)

	msg := &dsc3.Id{Value: ulid.Make().Bytes()}

	if err := v.Validate(msg); err != nil {
		assert.NoError(t, err)
	}

	if err = v.Validate(msg); err != nil {
		t.Logf("validation failed: %s", err)
	} else {
		t.Log("validation succeeded")
	}
}

func TestId2Validation(t *testing.T) {
	v, newErr := protovalidate.New()
	require.NoError(t, newErr)

	err := v.Validate(&dsc3.Id{Value: make([]byte, 17)})
	assert.Error(t, err)

	var valErr *protovalidate.ValidationError
	assert.ErrorAs(t, err, &valErr)

	if ok := errors.As(err, &valErr); ok {
		pb := valErr.ToProto()
		t.Log(pb.GetViolations())
	}
}
