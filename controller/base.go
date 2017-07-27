package controller

import (
	"encoding/json"
	"errors"
	"github.com/yinhui87/iris-component/api-response"
	"gopkg.in/kataras/iris.v6"
	"io/ioutil"
	"runtime"
	"strconv"
)

// outputJson
func Output(ctx *iris.Context, data interface{}, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctx.RequestHeader("token")
	}
	resp.Assign(response.RET_NORMAL, response.ERROR_NO_ERROR, "", token, data)
	ctx.JSON(iris.StatusOK, resp)
}

// success
func Success(ctx *iris.Context, info string, data interface{}, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctx.RequestHeader("token")
	}
	resp.Assign(response.RET_NORMAL, response.ERROR_NO_ERROR, info, token, data)
	ctx.JSON(iris.StatusOK, resp)
}

func logErrToConsole(ctx *iris.Context, logInfo string) {
	if logInfo != "" {
		funcName, file, line, ok := runtime.Caller(1)
		if ok {
			ctx.Log(iris.ProdMode, "%s\nFunc name: %s\nFile: %s[%d]\n", logInfo, runtime.FuncForPC(funcName).Name(), file, line)
		} else {
			ctx.Log(iris.ProdMode, logInfo)
		}
	}
}

func InvalidParam(ctx *iris.Context, code int32, info string, data interface{}, logInfo string, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctx.RequestHeader("token")
	}
	logErrToConsole(ctx, logInfo)
	resp.Assign(response.RET_ERROR, code, info, token, data)
	ctx.JSON(iris.StatusOK, resp)
	panic(nil)
}

// outputJson
func Error(ctx *iris.Context, code int32, info string, data interface{}, logInfo string, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctx.RequestHeader("token")
	}
	logErrToConsole(ctx, logInfo)
	resp.Assign(response.RET_ERROR, code, info, token, data)
	panic(resp)
}

func NotFound(ctx *iris.Context, info string, data interface{}, toKen ...string) {
	resp := response.RespondData{}
	var token string
	if len(toKen) > 0 {
		token = toKen[0]
	} else {
		token = ctx.RequestHeader("token")
	}
	resp.Assign(response.RET_ERROR, response.ERROR_NORMAL_ERROR, info, token, data)
	ctx.JSON(iris.StatusNotFound, resp)
	panic(nil)
}

func GetLimitOffset(ctx *iris.Context, convertToOffset bool) (limit int64, offset int64) {
	var err error
	var page int64
	limit, err = strconv.ParseInt(string(ctx.FormValue("size")), 10, 64)
	if err != nil || limit < 1 {
		limit = 10
	}
	page, err = strconv.ParseInt(string(ctx.FormValue("page")), 10, 64)
	if err != nil || page < 1 {
		page = 1
	}
	if !convertToOffset {
		return limit, page
	}
	offset = PageToOffset(page, limit)
	return
}

func PageToOffset(page, limit int64) (offset int64) {
	// calculate page
	if page < 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	return
}

func UnmarshalJsonBody(ctx *iris.Context, item interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return errors.New("ReadBody error: " + err.Error())
	}
	if err = json.Unmarshal(body, item); err != nil {
		return errors.New("Unmarshal body error: " + err.Error())
	}
	return nil
}
