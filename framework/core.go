package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore()*Core{
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRoute(url, allHandlers); err != nil {
		log.Fatal("add route error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRoute(url, allHandlers); err != nil {
		log.Fatal("add route error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRoute(url, allHandlers); err != nil {
		log.Fatal("add route error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRoute(url, allHandlers); err != nil {
		log.Fatal("add route error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup{
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler{
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	ctx := NewContext(request, response)
	handlers := c.FindRouteByRequest(request)
	if handlers == nil{
		ctx.Json(404,"not found")
		return
	}
	ctx.SetHandlers(handlers)
	if err := ctx.Next(); err != nil{
		ctx.Json(500,"internal error")
		return
	}
}

func (c *Core) Use(middlewares ...ControllerHandler){
	c.middlewares = append(c.middlewares, middlewares...)
}