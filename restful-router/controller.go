package restful

import (
	"gopkg.in/kataras/iris.v6"
)

type Controller struct {
}

func (this *Controller) Index(ctx *iris.Context) {
	ctx.NotFound()
}

func (this *Controller) Store(ctx *iris.Context) {
	ctx.NotFound()
}

func (this *Controller) Show(ctx *iris.Context, id string) {
	ctx.NotFound()
}

func (this *Controller) Update(ctx *iris.Context, id string) {
	ctx.NotFound()
}

func (this *Controller) Destroy(ctx *iris.Context, id string) {
	ctx.NotFound()
}
