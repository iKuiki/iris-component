package response

import (
	"gopkg.in/kataras/iris.v6"
)

var ErrHandler = iris.HandlerFunc(func(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			resp, ok := err.(RespondData)
			if ok {
				ctx.JSON(iris.StatusInternalServerError, resp)
				ctx.StopExecution()
			} else {
				ctx.Log(iris.DevMode, "[%s]%s Panic: %#v", ctx.Method(), ctx.Path(), err)
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
