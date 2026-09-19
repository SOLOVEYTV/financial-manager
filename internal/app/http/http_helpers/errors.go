package http_helpers

const (
	CodeNotFound           = "NOT_FOUND"
	CodeInvalidRequest     = "INVALID_REQUEST"
	CodeInternalError      = "INTERNAL_ERROR"
	CodeDecodeFailed       = "DECODE_FAILED"
	CodeUnauthorized       = "UNAUTHORIZED"
	CodeInvalidCredentials = "INVALID_CREDENTIALS"
	CodeForbidden          = "FORBIDDEN"
	CodeConflict           = "CONFLICT"
	CodeValidation         = "VALIDATION_ERROR"
)

const (
	MsgInvalidJSON   = "invalid JSON body"
	MsgInternalError = "internal server error"
	MsgUnauthorized  = "unauthorized"
)

const (
	ErrTxtEmpty          = "can not be blank"
	ErrTxtInvalidFormat  = "invalid format"
	ErrTxtUnknown        = "unknown value"
	ErrTxtTooShort       = "too short"
	ErrTxtTooLong        = "too long"
	ErrTxtMustBePositive = "must be positive"
	ErrTxtZeroValue      = "zero value"
)
