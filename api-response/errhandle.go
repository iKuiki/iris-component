package response

import (
	"gopkg.in/kataras/iris.v6"
	"runtime"
)

var ErrHandler = iris.HandlerFunc(func(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			resp, ok := err.(RespondData)
			if ok {
				ctx.JSON(iris.StatusInternalServerError, resp)
				ctx.StopExecution()
			} else {
				funcName, file, line, ok := runtime.Caller(3)
				if ok {
					ctx.Log(iris.DevMode, "[%s]%s Panic: %#v\nFunc name: %s\nFile: %s[%d]\n", ctx.Method(), ctx.Path(), err, runtime.FuncForPC(funcName).Name(), file, line)
				} else {
					ctx.Log(iris.DevMode, "[%s]%s Panic: %#v\n", ctx.Method(), ctx.Path(), err)
				}
				ctx.Panic()
			}

			//ctx.Panic just sends  http status 500 by default, but you can change it by: iris.OnPanic(func( c *iris.Context){})
		}
	}()
	ctx.Next()
})

// New restores the server on internal server errors (panics)
// returns the middleware
//
// is here for compatiblity
func NewErrHandler() iris.HandlerFunc {
	return ErrHandler
}
