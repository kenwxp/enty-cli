package util

import "strconv"

type MessageError struct {
	errorCode int
	errorType string
	msg       string
}

func (e *MessageError) Error() string {
	errJson := "[ERROR MESSAGE: errorCode: " + strconv.Itoa(e.errorCode) +
		", errorType: " + e.errorType + ", message: " + e.msg + "]"
	errString := string(errJson)
	return errString
}
func (e *MessageError) SetMsg(msg string) {
	e.msg = msg
}
func (e *MessageError) getMsg() string {
	return e.msg
}
func NewMsgError(code int, msg string) *MessageError {
	err := MessageError{
		errorCode: code,
		msg:       msg,
	}
	switch err.errorCode {
	case 0:
		err.errorType = "math Error"
	case 1:
		err.errorType = "rpc-link Error"
	case 2:
		err.errorType = "dataBase Error"
	case 3:
		err.errorType = "read-json Error"
	case 4:
		err.errorType = "type-trans Error"
	case 5:
		err.errorType = "server Error"
	case 9:
		err.errorType = "other Error"
	case 99:
		err.errorType = "unknown Error"

	}
	return &err
}
