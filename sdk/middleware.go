package sdk

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *App) wrap(h AppHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		if err := h(w, r, ps); err != nil {
			a.handleError(w, r, err)
		}
	}
}

// handleError is a method to handle errors in the App struct.
func (a *App) handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error: %v", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
