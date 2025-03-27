package sdk

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *App) Group(prefix string) *App {
	return &App{
		router:     a.router,
		middleware: a.middleware,
		pool:       a.pool,
		prefix:     a.prefix + prefix,
	}
}

func (a *App) fullPath(path string) string {
	return a.prefix + path
}

func (a *App) wrapHandler(h HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := a.pool.Get().(*Context)
		ctx.Writer = w
		ctx.Request = r
		ctx.Params = ps

		for _, mw := range a.middleware {
			mw(ctx)
		}
		h(ctx)
		a.pool.Put(ctx)
	}
}

func (a *App) Get(path string, h HandlerFunc)    { a.router.GET(a.fullPath(path), a.wrapHandler(h)) }
func (a *App) Post(path string, h HandlerFunc)   { a.router.POST(a.fullPath(path), a.wrapHandler(h)) }
func (a *App) Put(path string, h HandlerFunc)    { a.router.PUT(a.fullPath(path), a.wrapHandler(h)) }
func (a *App) Delete(path string, h HandlerFunc) { a.router.DELETE(a.fullPath(path), a.wrapHandler(h)) }
func (a *App) Patch(path string, h HandlerFunc)  { a.router.PATCH(a.fullPath(path), a.wrapHandler(h)) }
func (a *App) Options(path string, h HandlerFunc) {
	a.router.OPTIONS(a.fullPath(path), a.wrapHandler(h))
}
func (a *App) Head(path string, h HandlerFunc) { a.router.HEAD(a.fullPath(path), a.wrapHandler(h)) }
