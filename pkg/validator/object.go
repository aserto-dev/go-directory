package validator

import (
	dsc3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	dsr3 "github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
	dsw3 "github.com/aserto-dev/go-directory/aserto/directory/writer/v3"
)

func Object(msg *dsc3.Object) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetId()); err != nil {
		return err
	}
	if err := DisplayName(msg.GetDisplayName()); err != nil {
		return err
	}
	if err := Etag(msg.GetEtag()); err != nil {
		return err
	}
	return nil
}

func ObjectIdentifier(msg *dsc3.ObjectIdentifier) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetObjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetObjectId()); err != nil {
		return err
	}
	return nil
}

func GetObjectRequest(msg *dsr3.GetObjectRequest) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetObjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetObjectId()); err != nil {
		return err
	}
	if err := PaginationRequest(msg.GetPage()); err != nil {
		return err
	}
	return nil
}

func GetObjectsRequest(msg *dsr3.GetObjectsRequest) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetObjectType()); err != nil {
		return err
	}
	if err := PaginationRequest(msg.GetPage()); err != nil {
		return err
	}
	return nil
}

func GetObjectManyRequest(msg *dsr3.GetObjectManyRequest) error {
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

func SetObjectRequest(msg *dsw3.SetObjectRequest) error {
	if msg == nil || msg.GetObject() == nil {
		return nil
	}
	if err := Object(msg.GetObject()); err != nil {
		return err
	}
	return nil
}

func DeleteObjectRequest(msg *dsw3.DeleteObjectRequest) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetObjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetObjectId()); err != nil {
		return err
	}
	return nil
}
