package validator

import (
	dsr "github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
	dsw "github.com/aserto-dev/go-directory/aserto/directory/writer/v3"
)

func GetManifestRequest(msg *dsr.GetManifestRequest) error {
	return nil
}

func SetManifestRequest(msg *dsw.SetManifestRequest) error {
	return nil
}

func DeleteManifestRequest(msg *dsw.DeleteManifestRequest) error {
	return nil
}

// func Metadata(msg *dsm3.Metadata) error {
// 	return nil
// }

// var ErrBodyDataSize = errors.New("data size exceeds max chunk size of 65536 bytes")

// func Body(msg *dsm3.Body) error {
// 	if msg == nil {
// 		return nil
// 	}

// 	if len(msg.GetData()) > model.MaxChunkSizeBytes {
// 		return ErrBodyDataSize
// 	}

// 	return nil
// }
