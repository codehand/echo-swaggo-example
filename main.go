package main

import (
	"github.com/codehand/echo-swaggo-example/api"
	"github.com/labstack/echo/v4"
)

// @title Swagger example API
// @version 1.0
// @description This is a swagger server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1
// @BasePath /
// @Schemes http
func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/example", api.GetExamples)
	e.POST("/example", api.PostExample)
	e.GET("/example/:example_id", api.GetExample)
	e.PUT("/example/:example_id", api.PutExample)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// vi:syntax=go
