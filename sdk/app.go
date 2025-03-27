package sdk

import (
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context)

type App struct {
	router     *httprouter.Router
	middleware []HandlerFunc
	pool       sync.Pool
	prefix     string
}

func New() *App {
	r := httprouter.New()
	app := &App{
		router: r,
		pool: sync.Pool{
			New: func() any {
				return &Context{}
			},
		},
	}
	return app
}

func (a *App) Use(mw HandlerFunc) {
	a.middleware = append(a.middleware, mw)
}

func (a *App) Listen(addr string) error {
	return http.ListenAndServe(addr, a.router)
}
