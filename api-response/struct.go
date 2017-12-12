package response

// RetStatus 返回状态
type RetStatus int32

const (
	// RetError 未按预期执行，有错误发生
	RetError RetStatus = -1
	// RetNormal 正常执行
	RetNormal RetStatus = 0
	// RetInvalidParam 因输入错误而无法执行
	RetInvalidParam RetStatus = 1
	// RetNeedlogin 因未登录而无法执行
	RetNeedlogin RetStatus = 2
)

// RespCode 若发生错误，错误的详细代码
type RespCode int32

const (
	// ErrorNoError 无错误
	ErrorNoError RespCode = 0
	// ErrorNormalError 一般错误
	ErrorNormalError RespCode = -1
)

// RespondData API返回结果的框架格式
type RespondData struct {
	// Status 执行结果
	Ret RetStatus `json:"ret"`
	// Code 错误代码，若无错误发生则未0(RetCodeNoError)
	Code RespCode `json:"code"`
	// Info 附加信息
	Info string `json:"info"`
	// Data 若有数据返回，则在此对象内
	Data interface{} `json:"data"`
	// Token 下次交互使用的token
	Token string `json:"token"`
}

// Assign 将参数一次性赋值到返回数据中
func (r *RespondData) Assign(ret RetStatus, code RespCode, info, token string, data interface{}) {
	r.Ret = ret
	r.Code = code
	r.Info = info
	r.Data = data
	r.Token = token
}
