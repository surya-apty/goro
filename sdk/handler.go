package sdk

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AppHandler func(http.ResponseWriter, *http.Request, httprouter.Params) error
