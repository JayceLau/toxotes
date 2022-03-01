package terror

import "fmt"

// Err 通用固定错误
type Err struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Err) GetCode() uint32 {
	return e.Code
}

func (e *Err) GetMsg() string {
	return e.Msg
}

func (e *Err) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Msg)
}

func NewErr(code uint32, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}
func NewSuccess(msg string) *Err {
	return NewErr(OK, msg)
}

func NewErrCode(code uint32) *Err {
	return &Err{
		Code: code,
		Msg:  MapErrMsg(code),
	}
}

func NewErrMsg(msg string) *Err {
	return &Err{
		Code: ErrCommon,
		Msg:  msg,
	}
}
