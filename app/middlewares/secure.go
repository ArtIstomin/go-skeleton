package middleware

import (
	"github.com/iris-contrib/middleware/secure"
)

var helmet = secure.New(secure.Options{
	SSLRedirect:        true,
	FrameDeny:          true,
	ContentTypeNosniff: true,
	BrowserXSSFilter:   true,
	IsDevelopment:      isDevelopment,
})
