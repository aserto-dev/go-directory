package validate_test

import (
	"testing"

	dsc3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	"github.com/bufbuild/protovalidate-go"
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
