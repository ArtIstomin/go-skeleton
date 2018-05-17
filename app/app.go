package app

import (
	middleware "github.com/artistomin/friend4me/app/middlewares"
	route "github.com/artistomin/friend4me/app/routes"
	"github.com/gobuffalo/envy"
	"github.com/kataras/iris"
)

//Start application
func Start() {
	port := (":" + envy.Get("PORT", "3000"))
	app := iris.New()

	middleware.Bind(app)
	route.Bind(app)

	app.Run(
		iris.Addr(port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
