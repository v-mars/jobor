package bind

import "github.com/cloudwego/hertz/pkg/app/server/binding"

type BindError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *BindError) Error() string {
	if e.Msg != "" {
		return e.ErrType + ": 表达式路径=" + e.FailField + ", 参数错误原因=" + e.Msg
	}
	return e.ErrType + ": 表达式路径=" + e.FailField + ", 参数错误原因=无效的"
}

type ValidateError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *ValidateError) Error() string {
	if e.Msg != "" {
		return e.ErrType + ": 表达式路径=" + e.FailField + ", 验证错误原因=" + e.Msg
	}
	return e.ErrType + ": 表达式路径=" + e.FailField + ", 验证错误原因=无效的" // cause=invalid
}

func Init() {
	CustomBindErrFunc := func(failField, msg string) error {
		err := BindError{
			ErrType:   "参数绑定错误",
			FailField: "[绑定失败字段]: " + failField,
			Msg:       msg,
			//Msg:       "[绑定错误信息]: " + msg,
		}

		return &err
	}

	CustomValidateErrFunc := func(failField, msg string) error {
		err := ValidateError{
			ErrType:   "参数验证错误",
			FailField: "[验证失败字段]: " + failField,
			Msg:       msg,
			//Msg:       "[验证错误信息]: " + msg,
		}

		return &err
	}

	binding.SetErrorFactory(CustomBindErrFunc, CustomValidateErrFunc)

	// 默认false，全局生效
	binding.SetLooseZeroMode(true)

	// 使用标准库
	binding.UseStdJSONUnmarshaler()

	// 使用gjson
	//binding.UseGJSONUnmarshaler()

	// 使用第三方json unmarshal方法
	//binding.UseThirdPartyJSONUnmarshaler()
}
