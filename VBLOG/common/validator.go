package common

import "github.com/go-playground/validator/v10"

// 使用统计的校验函数-抽取到公共方法类
func Validate(obj any) error {
	v := validator.New()

	// 校验规则通过写struct tag来进定义
	return v.Struct(obj)
}
