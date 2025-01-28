package validator

import (
	dsc3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	dsr3 "github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
	dsw3 "github.com/aserto-dev/go-directory/aserto/directory/writer/v3"
)

func Relation(msg *dsc3.Relation) error {
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
	if err := TypeIdentifier(msg.GetSubjectRelation()); err != nil {
		return err
	}
	if err := Etag(msg.GetEtag()); err != nil {
		return err
	}
	return nil
}

func RelationIdentifier(msg *dsc3.RelationIdentifier) error {
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
	if err := TypeIdentifier(msg.GetSubjectRelation()); err != nil {
		return err
	}
	return nil
}

func GetRelationRequest(msg *dsr3.GetRelationRequest) error {
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
	if err := TypeIdentifier(msg.GetSubjectRelation()); err != nil {
		return err
	}
	return nil
}

func GetRelationsRequest(msg *dsr3.GetRelationsRequest) error {
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
	if err := TypeIdentifier(msg.GetSubjectRelation()); err != nil {
		return err
	}
	if err := PaginationRequest(msg.GetPage()); err != nil {
		return err
	}
	return nil
}

func SetRelationRequest(msg *dsw3.SetRelationRequest) error {
	if msg == nil || msg.GetRelation() == nil {
		return nil
	}
	if err := Relation(msg.GetRelation()); err != nil {
		return err
	}
	return nil
}

func DeleteRelationRequest(msg *dsw3.DeleteRelationRequest) error {
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
	if err := TypeIdentifier(msg.GetSubjectRelation()); err != nil {
		return err
	}
	return nil
}
