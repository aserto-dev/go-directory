package derr

import (
	"net/http"

	cerr "github.com/aserto-dev/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrUnknown                       = cerr.NewAsertoError("E20000", codes.Internal, http.StatusInternalServerError, "an unknown error has occurred")
	ErrInvalidJSON                   = cerr.NewAsertoError("E20001", codes.Internal, http.StatusInternalServerError, "invalid json")
	ErrNotFound                      = cerr.NewAsertoError("E20002", codes.NotFound, http.StatusNotFound, "not found")
	ErrAlreadyExists                 = cerr.NewAsertoError("E20003", codes.AlreadyExists, http.StatusConflict, "already exists")
	ErrInvalidObjectType             = cerr.NewAsertoError("E20004", codes.InvalidArgument, http.StatusBadRequest, "invalid object type")
	ErrInUse                         = cerr.NewAsertoError("E20005", codes.FailedPrecondition, http.StatusPreconditionFailed, "in use")
	ErrDeprecated                    = cerr.NewAsertoError("E20006", codes.Unimplemented, http.StatusNotImplemented, "deprecated")
	ErrInvalidRelationType           = cerr.NewAsertoError("E20007", codes.InvalidArgument, http.StatusBadRequest, "invalid relation type")
	ErrInvalidPermission             = cerr.NewAsertoError("E20008", codes.InvalidArgument, http.StatusBadRequest, "invalid permission")
	ErrInvalidObject                 = cerr.NewAsertoError("E20009", codes.InvalidArgument, http.StatusBadRequest, "invalid object")
	ErrInvalidRelation               = cerr.NewAsertoError("E20010", codes.InvalidArgument, http.StatusBadRequest, "invalid relation")
	ErrInvalidRelationTypePermission = cerr.NewAsertoError("E20011", codes.InvalidArgument, http.StatusBadRequest, "invalid relation type permission")
	ErrDirectoryStoreTenantNotFound  = cerr.NewAsertoError("E20012", codes.NotFound, http.StatusNotFound, "tenant store not found")
	ErrHashMismatch                  = cerr.NewAsertoError("E20013", codes.FailedPrecondition, http.StatusPreconditionFailed, "etag mismatch")
	ErrInvalidArgument               = cerr.NewAsertoError("E20015", codes.InvalidArgument, http.StatusBadRequest, "invalid argument")
	ErrInvalidID                     = cerr.NewAsertoError("E20016", codes.InvalidArgument, http.StatusBadRequest, "invalid ID")
	ErrAuthenticationFailed          = cerr.NewAsertoError("E20017", codes.FailedPrecondition, http.StatusUnauthorized, "authentication failed")
	ErrAuthorizationFailed           = cerr.NewAsertoError("E20018", codes.PermissionDenied, http.StatusUnauthorized, "authorization failed")
	ErrInvalidDecision               = cerr.NewAsertoError("E20019", codes.InvalidArgument, http.StatusBadRequest, "invalid decision")
	ErrInvalidTenantID               = cerr.NewAsertoError("E20020", codes.InvalidArgument, http.StatusBadRequest, "invalid tenant id")
	ErrInvalidTenantName             = cerr.NewAsertoError("E20021", codes.InvalidArgument, http.StatusBadRequest, "invalid tenant name")
	ErrDirectoryObjectNotFound       = cerr.NewAsertoError("E20022", codes.NotFound, http.StatusNotFound, "directory object not found")
	ErrNoTenantID                    = cerr.NewAsertoError("E20023", codes.InvalidArgument, http.StatusBadRequest, "no tenant id specified")
	ErrMaxDepthExceeded              = cerr.NewAsertoError("E20024", codes.Aborted, http.StatusInternalServerError, "graph depth exceeded")
	ErrObjectNotFound                = cerr.NewAsertoError("E20025", codes.NotFound, http.StatusNotFound, "object not found")
	ErrObjectTypeNotFound            = cerr.NewAsertoError("E20026", codes.NotFound, http.StatusNotFound, "object type not found")
	ErrObjectTypeAlreadyExists       = cerr.NewAsertoError("E20027", codes.AlreadyExists, http.StatusConflict, "object type already exists")
	ErrObjectTypeInUse               = cerr.NewAsertoError("E20028", codes.FailedPrecondition, http.StatusPreconditionFailed, "object type in use")
	ErrObjectInUse                   = cerr.NewAsertoError("E20029", codes.FailedPrecondition, http.StatusPreconditionFailed, "object is in use by one or more relations")
	ErrGroupNotFound                 = cerr.NewAsertoError("E20030", codes.NotFound, http.StatusNotFound, "group not found")
	ErrGroupAlreadyExists            = cerr.NewAsertoError("E20031", codes.AlreadyExists, http.StatusConflict, "group already exists")
	ErrSubjectNotFound               = cerr.NewAsertoError("E20032", codes.NotFound, http.StatusNotFound, "subject not found")
	ErrResourceNotFound              = cerr.NewAsertoError("E20033", codes.NotFound, http.StatusNotFound, "resource not found")
	ErrRelationTypeInUse             = cerr.NewAsertoError("E20034", codes.FailedPrecondition, http.StatusPreconditionFailed, "relation type in use")
	ErrRelationNotFound              = cerr.NewAsertoError("E20035", codes.NotFound, http.StatusNotFound, "relation not found")
	ErrRelationTypeNotFound          = cerr.NewAsertoError("E20036", codes.NotFound, http.StatusNotFound, "relation type not found")
	ErrRelationTypeAlreadyExists     = cerr.NewAsertoError("E20037", codes.AlreadyExists, http.StatusConflict, "relation type already exists")
	ErrPermissionAlreadyExists       = cerr.NewAsertoError("E20038", codes.AlreadyExists, http.StatusConflict, "permission already exists")
	ErrPermissionNotFound            = cerr.NewAsertoError("E20039", codes.NotFound, http.StatusNotFound, "permission not found")
	ErrPermissionInUse               = cerr.NewAsertoError("E20040", codes.FailedPrecondition, http.StatusPreconditionFailed, "permission in use by one or more relation types")
	ErrInvalidCursor                 = cerr.NewAsertoError("E20041", codes.InvalidArgument, http.StatusBadRequest, "invalid argument pagination cursor")
	ErrNotASubject                   = cerr.NewAsertoError("E20042", codes.InvalidArgument, http.StatusBadRequest, "invalid argument not a subject")
	ErrInvalidObjectTypeIdentifier   = cerr.NewAsertoError("E20043", codes.InvalidArgument, http.StatusBadRequest, "invalid argument object_type_identifier")
	ErrInvalidPermissionIdentifier   = cerr.NewAsertoError("E20044", codes.InvalidArgument, http.StatusBadRequest, "invalid argument permission_identifier")
	ErrInvalidRelationTypeIdentifier = cerr.NewAsertoError("E20045", codes.InvalidArgument, http.StatusBadRequest, "invalid argument relation_type_identifier")
	ErrInvalidObjectIdentifier       = cerr.NewAsertoError("E20046", codes.InvalidArgument, http.StatusBadRequest, "invalid argument object_identifier")
	ErrInvalidRelationIdentifier     = cerr.NewAsertoError("E20047", codes.InvalidArgument, http.StatusBadRequest, "invalid argument relation_identifier")
	ErrProtoValidate                 = cerr.NewAsertoError("E20050", codes.InvalidArgument, http.StatusBadRequest, "")
	ErrGraphDirectionality           = cerr.NewAsertoError("E20052", codes.InvalidArgument, http.StatusPreconditionFailed, "unable to determine graph directionality")
	ErrUnknownOpCode                 = cerr.NewAsertoError("E20053", codes.InvalidArgument, http.StatusPreconditionFailed, "unknown import OpCode value")
)
