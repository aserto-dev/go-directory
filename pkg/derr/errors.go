package derr

import (
	"net/http"

	cerr "github.com/aserto-dev/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrUnknown                       = newErr("E20000", codes.Internal, http.StatusInternalServerError, "an unknown error has occurred")
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
	ErrInvalidArgument               = newErr("E20015", codes.InvalidArgument, http.StatusBadRequest, "invalid argument")
	ErrInvalidID                     = newErr("E20016", codes.InvalidArgument, http.StatusBadRequest, "invalid ID")
	ErrAuthenticationFailed          = newErr("E20017", codes.FailedPrecondition, http.StatusUnauthorized, "authentication failed")
	ErrAuthorizationFailed           = newErr("E20018", codes.PermissionDenied, http.StatusUnauthorized, "authorization failed")
	ErrInvalidDecision               = newErr("E20019", codes.InvalidArgument, http.StatusBadRequest, "invalid decision")
	ErrInvalidTenantID               = newErr("E20020", codes.InvalidArgument, http.StatusBadRequest, "invalid tenant id")
	ErrInvalidTenantName             = newErr("E20021", codes.InvalidArgument, http.StatusBadRequest, "invalid tenant name")
	ErrDirectoryObjectNotFound       = newErr("E20022", codes.NotFound, http.StatusNotFound, "directory object not found")
	ErrNoTenantID                    = newErr("E20023", codes.InvalidArgument, http.StatusBadRequest, "no tenant id specified")

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
	ErrInvalidCursor                 = ErrInvalidArgument.Msg("pagination cursor")
	ErrNotASubject                   = ErrInvalidArgument.Msg("not a subject")
	ErrInvalidObjectTypeIdentifier   = ErrInvalidArgument.Msg("object_type_identifier")
	ErrInvalidPermissionIdentifier   = ErrInvalidArgument.Msg("permission_identifier")
	ErrInvalidRelationTypeIdentifier = ErrInvalidArgument.Msg("relation_type_identifier")
	ErrInvalidObjectIdentifier       = ErrInvalidArgument.Msg("object_identifier")
	ErrInvalidRelationIdentifier     = ErrInvalidArgument.Msg("relation_identifier")
	ErrInvalidObjectTypeID           = ErrInvalidArgument.Msg("object_type_id")
	ErrInvalidRelationTypeID         = ErrInvalidArgument.Msg("relation_type_id")
	ErrInvalidPermissionID           = ErrInvalidArgument.Msg("permission_id")
	ErrInvalidObjectID               = ErrInvalidArgument.Msg("object_id")
)

func newErr(code string, statusCode codes.Code, httpCode int, msg string) *cerr.AsertoError {
	return &cerr.AsertoError{Code: code, StatusCode: statusCode, Message: msg, HTTPCode: httpCode}
}
