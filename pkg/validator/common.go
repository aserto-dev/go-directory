package validator

import (
	"errors"

	dsc3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
)

var ErrPaginationSize = errors.New("pagination size must be >= 1 and <= 100")

func PaginationRequest(msg *dsc3.PaginationRequest) error {
	if msg == nil {
		return nil
	}

	if msg.GetSize() < 1 || msg.GetSize() > 100 {
		return ErrPaginationSize
	}

	return nil
}
