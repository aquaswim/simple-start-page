package error

import "fmt"

type AppError struct {
	Code        AppErrCode  `json:"code"`
	Message     string      `json:"message"`
	ErrorDetail interface{} `json:"errorDetail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func (a AppError) Error() string {
	return fmt.Sprintf("%d: %s", a.Code, a.Message)
}
