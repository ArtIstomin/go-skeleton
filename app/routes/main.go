package route

import (
	"github.com/kataras/iris"
)

// Bind routes to application
func Bind(app *iris.Application) {
	router := app.Party("/api")

	health(router)
}
