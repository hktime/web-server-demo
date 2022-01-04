package main

import (
	"net/http"
	"web-demo/framework"
)

func main()  {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
