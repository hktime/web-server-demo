package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)


type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context
	handler ControllerHandler

	hasTimeout bool
	writerMux *sync.Mutex
}

func NewContext(request *http.Request, response http.ResponseWriter)*Context{
	return &Context{
		request: request,
		responseWriter: response,
		ctx:            request.Context(),
		writerMux:      &sync.Mutex{},
	}
}

// base

func (ctx *Context) WriterMux() *sync.Mutex{
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request{
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter{
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
	return
}

func (ctx *Context) HasTimeout() bool{
	return ctx.hasTimeout
}
// base

// Context

func (ctx *Context) BaseContext()context.Context{
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{}{
	return ctx.BaseContext().Done()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool){
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Err() error{
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{}{
	return ctx.BaseContext().Value(key)
}
// context

// query

func (ctx *Context) QueryAll()map[string][]string{
	if ctx.request != nil{
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int)int{
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		leng := len(vals)
		if leng > 0{
			intVal, err := strconv.Atoi(vals[leng-1])
			if err != nil{
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		leng := len(vals)
		if leng > 0 {
			return vals[leng-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}
// end query

// post query

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		leng := len(vals)
		if leng > 0 {
			intVal, err := strconv.Atoi(vals[leng-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		leng := len(vals)
		if leng > 0 {
			return vals[leng-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}
// end post query

// response

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
// end response

