package response

type RET_STATUS int32

const (
	RET_ERROR     RET_STATUS = 0
	RET_NORMAL    RET_STATUS = 1
	RET_NEEDLOGIN RET_STATUS = 2
)

const (
	ERROR_NO_ERROR     int32 = 0
	ERROR_NORMAL_ERROR int32 = -1
)

type RespondData struct {
	Ret   RET_STATUS  `json:"ret"`
	Code  int32       `json:"code"`
	Info  string      `json:"info"`
	Data  interface{} `json:"data"`
	Token string      `json:"token"`
}

func (r *RespondData) Assign(ret RET_STATUS, code int32, info, token string, data interface{}) {
	r.Ret = ret
	r.Code = code
	r.Info = info
	r.Data = data
	r.Token = token
}
