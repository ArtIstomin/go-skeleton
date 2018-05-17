package route

import (
	"time"

	"github.com/kataras/iris"
)

var now = time.Now().UTC()

func health(router iris.Party) {
	router.Get("/health", func(ctx iris.Context) {
		uptime := time.Since(now).Round(time.Millisecond)
		ctx.JSON(iris.Map{
			"started": now.Format(time.RFC3339),
			"uptime":  uptime.String(),
		})
	})
}
