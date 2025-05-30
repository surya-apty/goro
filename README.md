github.com/surya-apty/goro

# Quickstart Example

```go
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
		c.JSON(200, map[string]string{"message": "Hello, surya!"})
	})

	app.Listen(":8080")
}

```

<img width="1410" alt="Screenshot 2025-04-07 at 10 28 55 PM" src="https://github.com/user-attachments/assets/9b672a0b-4e44-4437-9be1-746026a553e3" />

### http://localhost:8080/static

<img width="1434" alt="Screenshot 2025-04-07 at 10 24 27 PM" src="https://github.com/user-attachments/assets/aa3c90c1-9785-4f89-96c3-be14aa0d214b" />

