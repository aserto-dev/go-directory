package validator

import (
	dsc3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"

	dsr3 "github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
	dsw3 "github.com/aserto-dev/go-directory/aserto/directory/writer/v3"
)

func GetManifestRequest(msg *dsr3.GetManifestRequest) error {
	return nil
}

func SetManifestRequest(msg *dsw3.SetManifestRequest) error {
	return nil
}

func DeleteManifestRequest(msg *dsw3.DeleteManifestRequest) error {
	return nil
}

func Manifest(msg *dsc3.Manifest) error {
	return nil
}
