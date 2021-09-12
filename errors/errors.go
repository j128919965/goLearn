package errors

type BizError struct {
	code  int
	msg   string
	cause error
}

func (e *BizError) Error() string {
	return e.msg
}

var NotFound = &BizError{
	code:  404,
	msg:   "未找到该记录",
}

func New(message string,code int) error {
	return &BizError{
		msg:  message,
		code: code,
	}
}

func Wrap(cause error, msg string, code int) error {
	return &BizError{
		msg: msg,
		code: code,
		cause: cause,
	}
}