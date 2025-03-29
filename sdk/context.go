package sdk

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  httprouter.Params
}

func (c *Context) JSON(code int, data any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

func (c *Context) BindJSON(dest any) error {
	defer c.Request.Body.Close()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dest)
}

func (c *Context) HTML(code int, html string) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(html))
}
