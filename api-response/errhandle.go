package response

import (
	"github.com/kataras/iris"
	"runtime"
)

// ErrHandle 错误控制器
var ErrHandler = iris.Handler(func(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			resp, ok := err.(RespondData)
			if ok {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.JSON(resp)
				ctx.StopExecution()
			} else {
				funcName, file, line, ok := runtime.Caller(3)
				if ok {
					ctx.Application().Logger().Warnf("[%s]%s Panic: %#v\nFunc name: %s\nFile: %s[%d]\n", ctx.Method(), ctx.Path(), err, runtime.FuncForPC(funcName).Name(), file, line)
				} else {
					ctx.Application().Logger().Warnf("[%s]%s Panic: %#v\n", ctx.Method(), ctx.Path(), err)
				}
				ctx.StopExecution()
			}

			//ctx.Panic just sends  http status 500 by default, but you can change it by: iris.OnPanic(func( c *iris.Context){})
		}
	}()
	ctx.Next()
})

// NewErrHandler 返回错误容器中间件
func NewErrHandler() iris.Handler {
	return ErrHandler
}
