package response

import (
	"github.com/kataras/iris"
)

var ErrHandler = iris.HandlerFunc(func(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			resp, ok := err.(RespondData)
			if ok {
				ctx.JSON(iris.StatusInternalServerError, resp)
				ctx.StopExecution()
			} else {
				ctx.Log("Recovery from panic: %v\n", err)
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
