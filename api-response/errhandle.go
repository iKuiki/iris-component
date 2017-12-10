package response

import (
	"github.com/kataras/iris"
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
				ctx.Application().Logger().Errorf("[%s]%s Panic: %#v", ctx.Method(), ctx.Path(), err)
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
