package route

import (
	"github.com/gobuffalo/envy"
	"github.com/kataras/iris"
)

// Bind routes to application
func Bind(app *iris.Application) {
	prefix := envy.Get("API_PREFIX", "/api")
	router := app.Party(prefix)

	health(router)
}
