package response

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
)

type Code int64

func (p Code) String() string {
	switch p {
	case Err_NoLic:
		return "NoLic"
	case Err_BadRequest:
		return "BadRequest"
	case Err_AuthenticateFailed:
		return "AuthenticateFailed"
	case Err_Unauthenticated:
		return "Unauthenticated"
	case Err_Unauthorized:
		return "Unauthorized"
	case Err_ServerInternalErr:
		return "ServerInternalErr"
	case Err_ServerNotFound:
		return "ServerNotFound"
	case Err_RequestServerFail:
		return "RequestServerFail"
	case Err_ServerHandleFail:
		return "ServerHandleFail"
	case Err_BindAndValidateFail:
		return "BindAndValidateFail"
	case Err_Success:
		return "Success"
	case Err_ParamsErr:
		return "ParamsErr"
	case Err_AuthorizeFail:
		return "AuthorizeFail"
	case Err_RPCAuthSrvErr:
		return "RPCAuthSrvErr"
	case Err_AuthSrvErr:
		return "AuthSrvErr"
	case Err_RPCBlobSrvErr:
		return "RPCBlobSrvErr"
	case Err_BlobSrvErr:
		return "BlobSrvErr"
	case Err_RPCCarSrvErr:
		return "RPCCarSrvErr"
	case Err_CarSrvErr:
		return "CarSrvErr"
	case Err_RPCProfileSrvErr:
		return "RPCProfileSrvErr"
	case Err_ProfileSrvErr:
		return "ProfileSrvErr"
	case Err_RPCTripSrvErr:
		return "RPCTripSrvErr"
	case Err_TripSrvErr:
		return "TripSrvErr"
	case Err_RecordNotFound:
		return "RecordNotFound"
	case Err_RecordAlreadyExist:
		return "RecordAlreadyExist"
	case Err_DirtyData:
		return "DirtyData"
	}
	return "<UNSET>"
}

func ErrFromString(s string) (Code, error) {
	switch s {
	case "NoLic":
		return Err_NoLic, nil
	case "BadRequest":
		return Err_BadRequest, nil
	case "Unauthenticated":
		return Err_Unauthenticated, nil
	case "Unauthorized":
		return Err_Unauthorized, nil
	case "ServerInternalErr":
		return Err_ServerInternalErr, nil
	case "ServerNotFound":
		return Err_ServerNotFound, nil
	case "RequestServerFail":
		return Err_RequestServerFail, nil
	case "ServerHandleFail":
		return Err_ServerHandleFail, nil
	case "BindAndValidateFail":
		return Err_BindAndValidateFail, nil
	case "Success":
		return Err_Success, nil
	case "ParamsErr":
		return Err_ParamsErr, nil
	case "AuthorizeFail":
		return Err_AuthorizeFail, nil
	case "RPCAuthSrvErr":
		return Err_RPCAuthSrvErr, nil
	case "AuthSrvErr":
		return Err_AuthSrvErr, nil
	case "RPCBlobSrvErr":
		return Err_RPCBlobSrvErr, nil
	case "BlobSrvErr":
		return Err_BlobSrvErr, nil
	case "RPCCarSrvErr":
		return Err_RPCCarSrvErr, nil
	case "CarSrvErr":
		return Err_CarSrvErr, nil
	case "RPCProfileSrvErr":
		return Err_RPCProfileSrvErr, nil
	case "ProfileSrvErr":
		return Err_ProfileSrvErr, nil
	case "RPCTripSrvErr":
		return Err_RPCTripSrvErr, nil
	case "TripSrvErr":
		return Err_TripSrvErr, nil
	case "RecordNotFound":
		return Err_RecordNotFound, nil
	case "RecordAlreadyExist":
		return Err_RecordAlreadyExist, nil
	case "DirtyData":
		return Err_DirtyData, nil
	}
	return Code(0), fmt.Errorf("not a valid Code string")
}

func ErrPtr(v Code) *Code { return &v }
func (p *Code) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = Code(result.Int64)
	return
}

func (p *Code) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type StateCode struct {
	Name string `json:"name"`
	Code Code   `json:"code"`
}

var StateCodeList = []StateCode{
	/*{Name: Err_Success.String(), Code: Err_Success},
	{Name: Err_Unauthenticated.String(), Code: Err_Unauthenticated},
	{Name: Err_AuthenticateFailed.String(), Code: Err_AuthenticateFailed},
	{Name: Err_Unauthorized.String(), Code: Err_Unauthorized},
	{Name: Err_AuthorizeFail.String(), Code: Err_AuthorizeFail},
	{Name: Err_NoLic.String(), Code: Err_NoLic},
	{Name: Err_BadRequest.String(), Code: Err_BadRequest},
	{Name: Err_ParamsErr.String(), Code: Err_ParamsErr},
	{Name: Err_ServerNotFound.String(), Code: Err_ServerNotFound},
	{Name: Err_RequestServerFail.String(), Code: Err_RequestServerFail},
	{Name: Err_BindAndValidateFail.String(), Code: Err_BindAndValidateFail},
	{Name: Err_ServerHandleFail.String(), Code: Err_ServerHandleFail},
	{Name: Err_ServerInternalErr.String(), Code: Err_ServerInternalErr},
	{Name: Err_RPCAuthSrvErr.String(), Code: Err_RPCAuthSrvErr},
	{Name: Err_AuthSrvErr.String(), Code: Err_AuthSrvErr},
	{Name: Err_RPCBlobSrvErr.String(), Code: Err_RPCBlobSrvErr},
	{Name: Err_BlobSrvErr.String(), Code: Err_BlobSrvErr},
	{Name: Err_RecordNotFound.String(), Code: Err_RecordNotFound},
	{Name: Err_RecordAlreadyExist.String(), Code: Err_RecordAlreadyExist},*/
}

const (
	// 客户端错误 4xxxx
	// 服务端错误 5xxxx
	// 服务逻辑错误 6xxxx
	// 权限错误 1xxxx
	// http 状态码：20000
	//Err_BadRequest          Code = 10001
	//Err_Unauthorized        Code = 10002

	Err_Success             = 20000
	Err_Unauthenticated     = 10000
	Err_AuthenticateFailed  = 10001
	Err_Unauthorized        = 10002
	Err_AuthorizeFail       = 10003
	Err_NoLic               = 10004
	Err_BadRequest          = 40000
	Err_ParamsErr           = 40001
	Err_ServerNotFound      = 40003
	Err_RequestServerFail   = 40004
	Err_BindAndValidateFail = 40006
	Err_ServerHandleFail    = 50001
	Err_ServerInternalErr   = 50002
	Err_RPCAuthSrvErr       = 50003
	Err_AuthSrvErr          = 50004
	Err_RPCBlobSrvErr       = 50005
	Err_BlobSrvErr          = 50006
	Err_RPCCarSrvErr        = 60000
	Err_CarSrvErr           = 60001
	Err_RPCProfileSrvErr    = 60002
	Err_ProfileSrvErr       = 60003
	Err_RPCTripSrvErr       = 60004
	Err_TripSrvErr          = 60005
	Err_RecordNotFound      = 70000
	Err_RecordAlreadyExist  = 70001
	Err_DirtyData           = 70003
	Err_DataNotFoundErr     = 40002

	// ErrTaskExist 任务名已存在
	ErrTaskExist = 10416
	// ErrTaskNotExist 任务不存在
	ErrTaskNotExist = 10417

	// ErrHostgroupExist 主机组已存在
	ErrHostgroupExist = 10418
	// ErrHostgroupNotExist 主机组不存在
	ErrHostgroupNotExist = 10419
	// ErrDelHostUseByOtherHG 正在被其他的主机组使用，不能删除
	ErrDelHostUseByOtherHG = 10420
	//ErrHostNotExist 主机不存在
	ErrHostNotExist = 10421

	// ErrCronExpr CronExpr表达式不规范
	ErrCronExpr = 10422

	// ErrTaskUseByOtherTask 别的任务依赖此任务，请先在其他的任务的父子任务中移除此任务
	ErrTaskUseByOtherTask = 10423

	// ErrDelHostGroupUseByTask 正在被其他的任务使用，不能删除
	ErrDelHostGroupUseByTask = 10424
	// ErrDelUserUseByOther // 请先删除此用户创建的主机组或者任务后再删除
	ErrDelUserUseByOther = 10425

	// ErrCtxDeadlineExceeded 调用超时
	ErrCtxDeadlineExceeded = 10600
	// ErrCtxCanceled 取消调用
	ErrCtxCanceled = 10601

	// ErrRPCUnauthenticated  密钥认证失败
	ErrRPCUnauthenticated = 10602
	// ErrRPCUnavailable 调用对端不可用
	ErrRPCUnavailable = 10603
	// ErrRPCUnknown 调用未知错误
	ErrRPCUnknown = 10604
	// ErrRPCNotValidHost  未发现worker
	ErrRPCNotValidHost = 10605
	// ErrRPCNotConnHost 未找到存活的worker
	ErrRPCNotConnHost = 10606
)

var (
	BadRequest          = NewErrNo(Err_BadRequest, "bad request")
	Unauthenticated     = NewErrNo(Err_Unauthenticated, "unauthenticated, please login")
	AuthenticateFailed  = NewErrNo(Err_AuthenticateFailed, "登录认证信息失败")
	ServerInternalErr   = NewErrNo(Err_ServerInternalErr, "server internal error")
	AuthorizeFail       = NewErrNo(Err_AuthorizeFail, "authorize failed")
	RPCAuthSrvErr       = NewErrNo(Err_RPCAuthSrvErr, "rpc auth service error")
	AuthSrvErr          = NewErrNo(Err_AuthSrvErr, "auth service error")
	RPCBlobSrvErr       = NewErrNo(Err_RPCBlobSrvErr, "rpc blob service error")
	BlobSrvErr          = NewErrNo(Err_BlobSrvErr, "blob service error")
	RPCCarSrvErr        = NewErrNo(Err_RPCCarSrvErr, "rpc car service error")
	CarSrvErr           = NewErrNo(Err_CarSrvErr, "car service error")
	RPCProfileSrvErr    = NewErrNo(Err_RPCProfileSrvErr, "rpc profile service error")
	ProfileSrvErr       = NewErrNo(Err_ProfileSrvErr, "profile service error")
	RPCTripSrvErr       = NewErrNo(Err_RPCTripSrvErr, "rpc trip service error")
	TripSrvErr          = NewErrNo(Err_TripSrvErr, "trip service error")
	RecordNotFound      = NewErrNo(Err_RecordNotFound, "record not found")
	RecordAlreadyExist  = NewErrNo(Err_RecordAlreadyExist, "record already exist")
	DirtyData           = NewErrNo(Err_DirtyData, "dirty data")
	Succeed             = NewErrNo(Err_Success, "操作成功")
	ServiceErr          = NewErrNo(Err_ServerInternalErr, "服务没有正常启动")
	ParamErr            = NewErrNo(Err_ParamsErr, "参数错误")
	Err                 = NewErrNo(Err_ServerInternalErr, "服务器内部错误")
	OidcErr             = NewErrNo(Err_Unauthenticated, "未从统一认证服务器获取用户信息")
	DataNotFoundErr     = NewErrNo(Err_DataNotFoundErr, "您访问的数据不存在")
	UnauthorizedErr     = NewErrNo(Err_Unauthorized, "您没有权限执行该操作")
	AuthorizeFailErr    = NewErrNo(Err_AuthorizeFail, "认证失败")
	NameAlreadyExistErr = NewErrNo(Err_ParamsErr, "您输入的名称已存在")
	JsonParamErr        = NewErrNo(Err_ParamsErr, "这不是一个Json格式数据")
	NodeNotFound        = NewErrNo(Err_ParamsErr, "节点不存在")
	// 企业微信
	WorkWechatResponseErr = NewErrNo(Err_ServerInternalErr, "企业微信返回信息异常")
)

type ErrNo struct {
	ErrCode uint32
	ErrMsg  string
	Status  string
}

// NewErrNo return ErrNo
func NewErrNo[T int | uint32 | int64](code T, msg string) ErrNo {
	var status = "success"
	if code != 200 && code != 20000 {
		status = "error"
	}
	return ErrNo{
		ErrCode: uint32(code),
		ErrMsg:  msg,
		Status:  status,
	}
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) String() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

func (e ErrNo) GetErrMsg() string {
	return e.ErrMsg
}

func (e ErrNo) GetErrCode() uint32 {
	return e.ErrCode
}

var msgCode = map[int]string{
	//ErrUnauthorized: "非法请求",
	//ErrBadRequest:   "请求参数错误",
	//
	//ErrUserPassword:  "用户或密码错误",
	//ErrUserForbid:    "禁止登陆",
	//ErrEmailExist:    "邮箱已经存在",
	//ErrUserNameExist: "用户名已存在",
	//ErrUserNotExist:  "用户不存在",

	ErrTaskExist:    "任务名已存在",
	ErrTaskNotExist: "任务不存在",

	ErrHostgroupExist:      "主机组已存在",
	ErrHostgroupNotExist:   "主机组不存在",
	ErrDelHostUseByOtherHG: "正在被其他的主机组使用，不能删除",
	ErrHostNotExist:        "主机不存在",

	ErrCronExpr: "CronExpr表达式不规范",

	ErrTaskUseByOtherTask:    "存在任务依赖此任务，请先在其他的任务的父子任务中移除此任务",
	ErrDelHostGroupUseByTask: "正在被其他的任务使用，不能删除",
	ErrDelUserUseByOther:     "请先删除此用户创建的主机组或者任务后再删除",

	//ErrInternalServer: "服务端错误",

	ErrCtxDeadlineExceeded: "调用超时",
	ErrCtxCanceled:         "取消调用",
	ErrRPCUnauthenticated:  "密钥认证失败",
	ErrRPCUnavailable:      "调用对端不可用",
	ErrRPCUnknown:          "调用未知错误",
	ErrRPCNotValidHost:     "未发现worker",
	ErrRPCNotConnHost:      "未找到存活的worker",
}

func GetMsg(code int) string {
	var (
		msg    string
		exists bool
	)

	if msg, exists = msgCode[code]; exists {
		return msg
	}
	return "unknown"
}

// GetMsgErr get error msg by code
func GetMsgErr(code int) error {
	msg := GetMsg(code)
	return errors.New(msg)
}
