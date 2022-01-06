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

func (c *Core) FindRouteNodeByRequest(request *http.Request) *node{
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	// 封装自定义context
	ctx := NewContext(request, response)

	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		// 如果没有找到，这里打印日志
		ctx.SetStatus(404).Json( "not found")
		return
	}
	ctx.SetHandlers(handlers)
	// 寻找路由
	node := c.FindRouteNodeByRequest(request)
	if node == nil{
		ctx.SetStatus(404).Json("not found")
		return
	}
	// 设置路由参数
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)
	if err := ctx.Next(); err != nil{
		ctx.SetStatus(500).Json("internal error")
		return
	}
}

func (c *Core) Use(middlewares ...ControllerHandler){
	c.middlewares = append(c.middlewares, middlewares...)
}