package error

type AppErrCode int

const (
	ErrCodeInvalidCredential AppErrCode = 3
	ErrCodeUnauthenticated   AppErrCode = 2
	ErrCodeValidationFailed  AppErrCode = 1
	ErrCodeInternal          AppErrCode = 99
)
