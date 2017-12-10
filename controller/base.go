package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/yinhui87/iris-component/api-response"
	"strconv"
)

// BaseController 基础控制器，提供通用方法
type BaseController struct {
	mvc.C
}

// Output 输出json数据
func (ctl *BaseController) Output(data interface{}, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctl.Ctx.GetHeader("token")
	}
	resp.Assign(response.RetNormal, response.ErrorNoError, "", token, data)
	ctl.Ctx.JSON(resp)
}

// Success 执行操作成功时调用，含info
func (ctl *BaseController) Success(info string, data interface{}, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctl.Ctx.GetHeader("token")
	}
	resp.Assign(response.RetNormal, response.ErrorNoError, info, token, data)
	ctl.Ctx.JSON(resp)
}

func (ctl *BaseController) logErrToConsole(logInfo string) {
	if logInfo != "" {
		ctl.Ctx.Application().Logger().Warnf("[%s]%s | %s", ctl.Ctx.Method(), ctl.Ctx.Path(), logInfo)
	}
}

// InvalidParam 请求错误，可以填写参数，但返回仍未200
func (ctl *BaseController) InvalidParam(code response.RespCode, info string, data interface{}, logInfo string, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctl.Ctx.GetHeader("token")
	}
	ctl.logErrToConsole(logInfo)
	resp.Assign(response.RetError, code, info, token, data)
	ctl.Ctx.JSON(resp)
	panic(nil)
}

// Error 出错时返回错误
func (ctl *BaseController) Error(code response.RespCode, info string, data interface{}, logInfo string, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctl.Ctx.GetHeader("token")
	}
	ctl.logErrToConsole(logInfo)
	resp.Assign(response.RetError, code, info, token, data)
	panic(resp)
}

// NotFound 找不到请求的对象时使用
func (ctl *BaseController) NotFound(info string, data interface{}, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctl.Ctx.GetHeader("token")
	}
	resp.Assign(response.RetError, response.ErrorNormalError, info, token, data)
	ctl.Ctx.StatusCode(iris.StatusNotFound)
	ctl.Ctx.JSON(resp)
	panic(nil)
}

// ObtainLimitOffset 获取请求中的size与page参数，并转化为sql的limit与offset
func (ctl *BaseController) ObtainLimitOffset(convertToOffset bool) (limit int64, offset int64) {
	var err error
	var page int64
	limit, err = strconv.ParseInt(string(ctl.Ctx.FormValue("size")), 10, 64)
	if err != nil || limit < 1 {
		limit = 10
	}
	page, err = strconv.ParseInt(string(ctl.Ctx.FormValue("page")), 10, 64)
	if err != nil || page < 1 {
		page = 1
	}
	if !convertToOffset {
		return limit, page
	}
	offset = PageToOffset(page, limit)
	return
}

// PageToOffset 将page与limit转化为offset
func PageToOffset(page, limit int64) (offset int64) {
	// calculate page
	if page < 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	return
}
