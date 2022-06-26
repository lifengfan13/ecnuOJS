package util

import "net/http"

type Error struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

const (
	RequestSuccess = iota
	CreateSuccess
	UpdateSuccess
	AcceptSuccess
	DeleteSuccess

	ServerInternalError

	RegisterUserFormatError
	RegisterEmailDuplicated
	RegisterTelephoneDuplicated

	LoginUserFormatError
	LoginUserNotExist
	LoginUserPasswordError

	AuthTokenError
	AuthTokenInvalid
)

var errMsgs map[int]*Error = map[int]*Error{
	ServerInternalError: {"server internal error", ServerInternalError},

	RegisterUserFormatError:     {"user register message format error", RegisterUserFormatError},
	RegisterEmailDuplicated:     {"user eamil is duplicated", RegisterEmailDuplicated},
	RegisterTelephoneDuplicated: {"user telephone is duplicated", RegisterTelephoneDuplicated},

	LoginUserFormatError:   {"user login message format error", LoginUserFormatError},
	LoginUserNotExist:      {"user is not exist", LoginUserNotExist},
	LoginUserPasswordError: {"user password error", LoginUserPasswordError},

	AuthTokenError:   {"call for API without auth token", AuthTokenError},
	AuthTokenInvalid: {"user token invalid", AuthTokenInvalid},
}

var errStatus map[int]int = map[int]int{
	ServerInternalError: http.StatusInternalServerError,

	RegisterUserFormatError:     http.StatusBadRequest,
	RegisterEmailDuplicated:     http.StatusUnauthorized,
	RegisterTelephoneDuplicated: http.StatusUnauthorized,

	LoginUserFormatError:   http.StatusBadRequest,
	LoginUserNotExist:      http.StatusNotFound,
	LoginUserPasswordError: http.StatusUnauthorized,

	AuthTokenError:   http.StatusUnauthorized,
	AuthTokenInvalid: http.StatusUnauthorized,
}

func ReturnErrMsg(code int) *Error {
	return errMsgs[code]
}

func ReturnErrStatus(code int) int {
	return errStatus[code]
}
