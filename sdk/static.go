package sdk

import (
	"net/http"
	"os"
	"path/filepath"
)

func (a *App) Static(route string, dir string) {
	fileServer := http.StripPrefix(route, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(dir, "index.html"))
			return
		}
		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	}))
	a.router.Handler("GET", route+"/*filepath", fileServer)
}
