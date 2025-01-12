package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发送异常消息
func SendFaild(c *gin.Context, err error) {
	// 中断请求
	defer c.Abort()

	// 断言用户有没有实现自定义异常也有可能是非自定义
	// 1. errors.New() 参数的标准异常实现, msg, err.Error()
	// 2. 业务自定义异常: ApiException
	// 3. 其他实现, 比如阿里云 SDK 也封装SDK Error, 也是一种自定义异常
	ae, ok := err.(*ApiException)

	// 自定义的异常返回
	if ok {
		c.JSON(ae.HttpCode, NewApiException(*ae.ErrorCode, ae.Message))
		return
	}

	// 非自定义异常
	// ae.Message 断言失败 ae为nil
	c.JSON(http.StatusInternalServerError, NewApiException(-1, err.Error()))
}

func NewUnauthorized() *ApiException {
	return NewApiException(http.StatusUnauthorized, "请认证")
}

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		ErrorCode: &code,
		HttpCode:  code,
		Message:   message,
	}
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
// type error interface {
// 	Error() string
// }

// 自定义异常实现
type ApiException struct {
	// 0表示正常， nil 表示状态未知, 业务异常吗
	ErrorCode *int `json:"error_code"`
	// http 状态码
	HttpCode int `json:"http_code"`
	// 具体的异常信息
	Message string `json:"message"`
}

// 实现error接口
func (e *ApiException) Error() string {
	return e.Message
}

// 设置http状态码
func (e *ApiException) SetHttpCode(code int) *ApiException {
	e.HttpCode = code
	return e
}

// fmt Stringer
func (e *ApiException) ToJSON() string {
	d, _ := json.Marshal(e)
	return string(d)
}

//----------自己写的

type Response struct {
	ErrorCode int    `json:error_code`
	Message   string `json:message`
}

func NewResponse(code int, message string) *Response {
	return &Response{
		ErrorCode: code,
		Message:   message,
	}
}
