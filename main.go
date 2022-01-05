package main

import (
	"net/http"
	"web-demo/framework"
	"web-demo/framework/middlerware"
)

func main()  {
	core := framework.NewCore()
	core.Use(
		middlerware.Recovery(),
		middlerware.Cost(),
		)
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
