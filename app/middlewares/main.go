package middleware

import (
	"github.com/gobuffalo/envy"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var env = envy.Get("ENV", "development")
var isDevelopment = (env == "development")

//Bind middlewares to application
func Bind(app *iris.Application) {
	app.Logger().SetLevel(envy.Get("LOG_LEVEL", "debug"))

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.Serve)
	app.Use(cors.Default())
}
