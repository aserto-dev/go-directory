package derr

import (
	"net/http"

	"github.com/aserto-dev/go-utils/cerr"
	"google.golang.org/grpc/codes"
)

var (
	ErrInvalidJSON                   = newErr("E20001", codes.Internal, http.StatusInternalServerError, "invalid json")
	ErrNotFound                      = newErr("E20002", codes.NotFound, http.StatusNotFound, "not found")
	ErrAlreadyExists                 = newErr("E20003", codes.AlreadyExists, http.StatusConflict, "already exists")
	ErrInvalidObjectType             = newErr("E20004", codes.Internal, http.StatusInternalServerError, "invalid object type")
	ErrObjectTypeInUse               = newErr("E20005", codes.FailedPrecondition, http.StatusPreconditionFailed, "object type in use")
	ErrDeprecated                    = newErr("E20006", codes.OK, http.StatusOK, "deprecated")
	ErrInvalidRelationType           = newErr("E20007", codes.Internal, http.StatusInternalServerError, "invalid relation type")
	ErrInvalidPermission             = newErr("E20008", codes.Internal, http.StatusInternalServerError, "invalid permission")
	ErrInvalidObject                 = newErr("E20009", codes.Internal, http.StatusInternalServerError, "invalid object")
	ErrInvalidRelation               = newErr("E20010", codes.Internal, http.StatusInternalServerError, "invalid relation")
	ErrInvalidRelationTypePermission = newErr("E20011", codes.Internal, http.StatusInternalServerError, "invalid relation type permission")
	ErrDirectoryStoreTenantNotFound  = newErr("E20012", codes.NotFound, http.StatusNotFound, "tenant store not found")
	ErrHashMismatch                  = newErr("E20013", codes.InvalidArgument, http.StatusBadRequest, "hash value mismatch")
	ErrPermissionInUse               = newErr("E20014", codes.FailedPrecondition, http.StatusPreconditionFailed, "permission in use by one or more relation types")

	ErrObjectNotFound                = ErrNotFound.Str("type", "object")
	ErrObjectTypeNotFound            = ErrNotFound.Str("type", "object type")
	ErrObjectTypeAlreadyExists       = ErrAlreadyExists.Str("type", "object type")
	ErrGroupNotFound                 = ErrNotFound.Str("type", "group")
	ErrGroupAlreadyExists            = ErrAlreadyExists.Str("type", "group")
	ErrSubjectNotFound               = ErrNotFound.Str("type", "subject")
	ErrResourceNotFound              = ErrNotFound.Str("type", "resource")
	ErrRelationNotFound              = ErrNotFound.Str("type", "relation")
	ErrRelationTypeNotFound          = ErrNotFound.Str("type", "relation type")
	ErrRelationTypeAlreadyExists     = ErrAlreadyExists.Str("type", "relation type")
	ErrPermissionAlreadyExists       = ErrAlreadyExists.Str("type", "permission")
	ErrPermissionNotFound            = ErrNotFound.Str("type", "permission")
	ErrInvalidCursor                 = cerr.ErrInvalidArgument.Msg("pagination cursor")
	ErrNotASubject                   = cerr.ErrInvalidArgument.Msg("not a subject")
	ErrInvalidObjectTypeIdentifier   = cerr.ErrInvalidArgument.Msg("object_type_identifier")
	ErrInvalidPermissionIdentifier   = cerr.ErrInvalidArgument.Msg("permission_identifier")
	ErrInvalidRelationTypeIdentifier = cerr.ErrInvalidArgument.Msg("relation_type_identifier")
	ErrInvalidObjectIdentifier       = cerr.ErrInvalidArgument.Msg("object_identifier")
	ErrInvalidRelationIdentifier     = cerr.ErrInvalidArgument.Msg("relation_identifier")
	ErrInvalidObjectTypeID           = cerr.ErrInvalidArgument.Msg("object_type_id")
	ErrInvalidRelationTypeID         = cerr.ErrInvalidArgument.Msg("relation_type_id")
	ErrInvalidPermissionID           = cerr.ErrInvalidArgument.Msg("permission_id")
	ErrInvalidObjectID               = cerr.ErrInvalidArgument.Msg("object_id")
)

func newErr(code string, statusCode codes.Code, httpCode int, msg string) *cerr.AsertoError {
	return &cerr.AsertoError{Code: code, StatusCode: statusCode, Message: msg, HttpCode: httpCode}
}
