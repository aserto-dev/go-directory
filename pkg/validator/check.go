package validator

import (
	"github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
)

func CheckRequest(msg *reader.CheckRequest) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldObjectType, msg.GetObjectType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldObjectID, msg.GetObjectId()); err != nil {
		return err
	}

	if err := TypeIdentifier(fieldRelation, msg.GetRelation()); err != nil {
		return err
	}

	if err := TypeIdentifier(fieldSubjectType, msg.GetSubjectType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldSubjectID, msg.GetSubjectId()); err != nil {
		return err
	}

	return nil
}

func ChecksRequest(msg *reader.ChecksRequest) error {
	if msg == nil {
		return nil
	}

	if msg.GetDefault() != nil {
		if err := CheckRequest(msg.GetDefault()); err != nil {
			return err
		}
	}

	for _, check := range msg.GetChecks() {
		if err := CheckRequest(check); err != nil {
			return err
		}
	}

	return nil
}

func GetGraphRequest(msg *reader.GetGraphRequest) error {
	if msg == nil {
		return nil
	}

	if err := TypeIdentifier(fieldObjectType, msg.GetObjectType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldObjectID, msg.GetObjectId()); err != nil {
		return err
	}

	if err := TypeIdentifier(fieldRelation, msg.GetRelation()); err != nil {
		return err
	}

	if err := TypeIdentifier(fieldSubjectType, msg.GetSubjectType()); err != nil {
		return err
	}

	if err := InstanceIdentifier(fieldSubjectID, msg.GetSubjectId()); err != nil {
		return err
	}

	if err := TypeIdentifier(fieldSubjectRelation, msg.GetSubjectRelation()); err != nil {
		return err
	}

	return nil
}
