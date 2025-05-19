package validator

import (
	"github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	"github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
	"github.com/aserto-dev/go-directory/aserto/directory/writer/v3"
)

func GetManifestRequest(msg *reader.GetManifestRequest) error {
	return nil
}

func SetManifestRequest(msg *writer.SetManifestRequest) error {
	return nil
}

func DeleteManifestRequest(msg *writer.DeleteManifestRequest) error {
	return nil
}

func Manifest(msg *common.Manifest) error {
	return nil
}

func GetModelRequest(msg *reader.GetModelRequest) error {
	return nil
}

func Model(msg *common.Model) error {
	return nil
}
