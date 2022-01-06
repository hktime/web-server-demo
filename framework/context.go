package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)


type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context

	hasTimeout bool
	writerMux *sync.Mutex

	handlers []ControllerHandler
	index int // 当前请求调用调用到哪个节点

	params map[string]string
}

func NewContext(request *http.Request, response http.ResponseWriter)*Context{
	return &Context{
		request: request,
		responseWriter: response,
		ctx:            request.Context(),
		writerMux:      &sync.Mutex{},
		index: -1,
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

func (ctx *Context) SetHandlers(handlers []ControllerHandler){
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string){
	ctx.params = params
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

// Next 核心函数，调用context的下一个函数
func (ctx *Context) Next() error{
	ctx.index++
	if ctx.index < len(ctx.handlers){
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}