package jsonerror

import (
	"fmt"
	"net/http"

	"entysquare/filer-backend/util"
)

// InvestorsError represents the "standard error response" in Matrix.
// http://matrix.org/docs/spec/client_server/r0.2.0.html#api-standards
type InvestorsError struct {
	ErrCode string `json:"errcode"`
	Err     string `json:"error"`
}

func (e InvestorsError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrCode, e.Err)
}

// InternalServerError returns a 500 Internal Server Error in a matrix-compliant
// format.
func InternalServerError() util.JSONResponse {
	return util.JSONResponse{
		Code: http.StatusInternalServerError,
		JSON: Unknown("Internal Server Error"),
	}
}

// Unknown is an unexpected error
func Unknown(msg string) *InvestorsError {
	return &InvestorsError{"M_UNKNOWN", msg}
}

// Forbidden is an error when the client tries to access a resource
// they are not allowed to access.
func Forbidden(msg string) *InvestorsError {
	return &InvestorsError{"M_FORBIDDEN", msg}
}

// BadJSON is an error when the client supplies malformed JSON.
func BadJSON(msg string) *InvestorsError {
	return &InvestorsError{"M_BAD_JSON", msg}
}

// BadRPC is an error when the client supplies malformed JSON.
func BadRPC(msg string) *InvestorsError {
	return &InvestorsError{"M_BAD_RPC", msg}
}

// NotJSON is an error when the client supplies something that is not JSON
// to a JSON endpoint.
func NotJSON(msg string) *InvestorsError {
	return &InvestorsError{"M_NOT_JSON", msg}
}

// NotFound is an error when the client tries to access an unknown resource.
func NotFound(msg string) *InvestorsError {
	return &InvestorsError{"M_NOT_FOUND", msg}
}

// MissingArgument is an error when the client tries to access a resource
// without providing an argument that is required.
func MissingArgument(msg string) *InvestorsError {
	return &InvestorsError{"M_MISSING_ARGUMENT", msg}
}

// InvalidArgumentValue is an error when the client tries to provide an
// invalid value for a valid argument
func InvalidArgumentValue(msg string) *InvestorsError {
	return &InvestorsError{"M_INVALID_ARGUMENT_VALUE", msg}
}

// MissingToken is an error when the client tries to access a resource which
// requires authentication without supplying credentials.
func MissingToken(msg string) *InvestorsError {
	return &InvestorsError{"M_MISSING_TOKEN", msg}
}

// UnknownToken is an error when the client tries to access a resource which
// requires authentication and supplies an unrecognised token
func UnknownToken(msg string) *InvestorsError {
	return &InvestorsError{"M_UNKNOWN_TOKEN", msg}
}

// WeakPassword is an error which is returned when the client tries to register
// using a weak password. http://matrix.org/docs/spec/client_server/r0.2.0.html#password-based
func WeakPassword(msg string) *InvestorsError {
	return &InvestorsError{"M_WEAK_PASSWORD", msg}
}

// InvalidUsername is an error returned when the client tries to register an
// invalid username
func InvalidUsername(msg string) *InvestorsError {
	return &InvestorsError{"M_INVALID_USERNAME", msg}
}

// UserInUse is an error returned when the client tries to register an
// username that already exists
func UserInUse(msg string) *InvestorsError {
	return &InvestorsError{"M_USER_IN_USE", msg}
}

// ASExclusive is an error returned when an application service tries to
// register an username that is outside of its registered namespace, or if a
// user attempts to register a username or room alias within an exclusive
// namespace.
func ASExclusive(msg string) *InvestorsError {
	return &InvestorsError{"M_EXCLUSIVE", msg}
}

// GuestAccessForbidden is an error which is returned when the client is
// forbidden from accessing a resource as a guest.
func GuestAccessForbidden(msg string) *InvestorsError {
	return &InvestorsError{"M_GUEST_ACCESS_FORBIDDEN", msg}
}

type IncompatibleRoomVersionError struct {
	RoomVersion string `json:"room_version"`
	Error       string `json:"error"`
	Code        string `json:"errcode"`
}

// UnsupportedRoomVersion is an error which is returned when the client
// requests a room with a version that is unsupported.
func UnsupportedRoomVersion(msg string) *InvestorsError {
	return &InvestorsError{"M_UNSUPPORTED_ROOM_VERSION", msg}
}

// LimitExceededError is a rate-limiting error.
type LimitExceededError struct {
	InvestorsError
	RetryAfterMS int64 `json:"retry_after_ms,omitempty"`
}

// LimitExceeded is an error when the client tries to send events too quickly.
func LimitExceeded(msg string, retryAfterMS int64) *LimitExceededError {
	return &LimitExceededError{
		InvestorsError: InvestorsError{"M_LIMIT_EXCEEDED", msg},
		RetryAfterMS:   retryAfterMS,
	}
}

// NotTrusted is an error which is returned when the client asks the server to
// proxy a request (e.g. 3PID association) to a server that isn't trusted
func NotTrusted(serverName string) *InvestorsError {
	return &InvestorsError{
		ErrCode: "M_SERVER_NOT_TRUSTED",
		Err:     fmt.Sprintf("Untrusted server '%s'", serverName),
	}
}
