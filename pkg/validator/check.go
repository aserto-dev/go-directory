package validator

import (
	dsr3 "github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
)

func CheckRequest(msg *dsr3.CheckRequest) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetObjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetObjectId()); err != nil {
		return err
	}
	if err := TypeIdentifier(msg.GetRelation()); err != nil {
		return err
	}
	if err := TypeIdentifier(msg.GetSubjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetSubjectId()); err != nil {
		return err
	}
	return nil
}

func ChecksRequest(msg *dsr3.ChecksRequest) error {
	if msg == nil {
		return nil
	}
	if msg.Default != nil {
		if err := CheckRequest(msg.Default); err != nil {
			return err
		}
	}
	for _, check := range msg.Checks {
		if err := CheckRequest(check); err != nil {
			return err
		}
	}
	return nil
}

// CheckPermissionRequest gets handled by CheckRequest, hence returns nil.
func CheckPermissionRequest(msg *dsr3.CheckPermissionRequest) error {
	return nil
}

// CheckRelationRequest gets handled by CheckRequest, hence returns nil.
func CheckRelationRequest(msg *dsr3.CheckRelationRequest) error {
	return nil
}

func GetGraphRequest(msg *dsr3.GetGraphRequest) error {
	if msg == nil {
		return nil
	}
	if err := TypeIdentifier(msg.GetObjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetObjectId()); err != nil {
		return err
	}
	if err := TypeIdentifier(msg.GetRelation()); err != nil {
		return err
	}
	if err := TypeIdentifier(msg.GetSubjectType()); err != nil {
		return err
	}
	if err := InstanceIdentifier(msg.GetSubjectId()); err != nil {
		return err
	}
	if err := TypeIdentifier((msg.SubjectRelation)); err != nil {
		return err
	}
	return nil
}
