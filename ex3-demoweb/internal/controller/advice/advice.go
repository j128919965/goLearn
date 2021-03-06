package advice

import "quezr.top/demoweb/internal/errors"

type ApiResponse struct {
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Success bool `json:"success"`
}

func Success(data interface{}) *ApiResponse {
	return &ApiResponse{Data: data,Success: true}
}

func BizFailure(err *errors.BizError) *ApiResponse {
	return &ApiResponse{Data: nil,Message: err.Msg ,Success: false}
}

func ErrFailure(err error) *ApiResponse {
	return &ApiResponse{Success: false,Message: err.Error()}
}
