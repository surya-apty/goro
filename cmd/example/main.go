package main

import (
	"github.com/surya-apty/goro/sdk"
)

func main() {
	app := sdk.New()

	app.Use(func(c *sdk.Context) {
		// Example middleware (noop)
	})

	app.Static("/static", "./public")

	app.Get("/", func(c *sdk.Context) {
		c.HTML(200, "<h1>Hello, Goro!</h1><p>Visit <a href='/static'>/static</a> for static files.</p>")
	})

	api := app.Group("/api/v1")

	api.Get("/hello", func(c *sdk.Context) {
		c.JSON(200, map[string]string{"message": "Hello, world!"})
	})

	app.Listen(":8080")
}
