package validator

import (
	"github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	"github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
	"github.com/aserto-dev/go-directory/aserto/directory/writer/v3"
)

func Object(msg *common.Object) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldType, msg.GetType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldID, msg.GetId()); err != nil {
		return err
	}

	if err := DisplayName(fieldDisplayName, msg.GetDisplayName()); err != nil {
		return err
	}

	if err := Etag(fieldETag, msg.GetEtag()); err != nil {
		return err
	}

	return nil
}

func ObjectIdentifier(msg *common.ObjectIdentifier) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldObjectType, msg.GetType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldObjectID, msg.GetId()); err != nil {
		return err
	}

	return nil
}

func GetObjectRequest(msg *reader.GetObjectRequest) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldObjectType, msg.GetObjectType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldObjectID, msg.GetObjectId()); err != nil {
		return err
	}

	if err := PaginationRequest(msg.GetPage()); err != nil {
		return err
	}

	return nil
}

func GetObjectsRequest(msg *reader.GetObjectsRequest) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldObjectType, msg.GetObjectType()); err != nil {
		return err
	}

	if err := PaginationRequest(msg.GetPage()); err != nil {
		return err
	}

	return nil
}

func GetObjectManyRequest(msg *reader.GetObjectManyRequest) error {
	if msg == nil || msg.GetParam() == nil || len(msg.GetParam()) == 0 {
		return nil
	}

	for _, oid := range msg.GetParam() {
		if err := ObjectIdentifier(oid); err != nil {
			return err
		}
	}

	return nil
}

func SetObjectRequest(msg *writer.SetObjectRequest) error {
	if msg == nil || msg.GetObject() == nil {
		return nil
	}

	if err := Object(msg.GetObject()); err != nil {
		return err
	}

	return nil
}

func DeleteObjectRequest(msg *writer.DeleteObjectRequest) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldObjectType, msg.GetObjectType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldObjectID, msg.GetObjectId()); err != nil {
		return err
	}

	return nil
}
