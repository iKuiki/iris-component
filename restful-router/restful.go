package restful

import (
	"gopkg.in/kataras/iris.v6"
)

type RESTfulController interface {
	Index(*iris.Context)
	Store(*iris.Context)
	Show(*iris.Context, string)
	Update(*iris.Context, string)
	Destroy(*iris.Context, string)
}

type IrisRouter interface {
	Get(string, ...iris.HandlerFunc) iris.RouteInfo
	Post(string, ...iris.HandlerFunc) iris.RouteInfo
	Put(string, ...iris.HandlerFunc) iris.RouteInfo
	Delete(string, ...iris.HandlerFunc) iris.RouteInfo
	Patch(string, ...iris.HandlerFunc) iris.RouteInfo
}

func adapt(function func(ctx *iris.Context, id string)) func(ctx *iris.Context) {
	return func(ctx *iris.Context) {
		id := ctx.Param("id")
		function(ctx, id)
	}
}

func Resource(router IrisRouter, controller RESTfulController) {
	router.Get("", controller.Index)
	router.Post("", controller.Store)
	router.Get("/:id", adapt(controller.Show))
	router.Put("/:id", adapt(controller.Update))
	router.Patch("/:id", adapt(controller.Update))
	router.Delete("/:id", adapt(controller.Destroy))
}
