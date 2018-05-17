package main

import (
	"github.com/artistomin/friend4me/app"
	"github.com/gobuffalo/envy"
)

func init() {
	envy.Load(".env")
	env := envy.Get("ENV", "development")

	switch env {
	case "staging":
		envy.Load(".env.staging")
	case "production":
		envy.Load(".env.production")
	}
}

func main() {
	app.Start()
}
