package main

import (
	"context"
	"github.com/hktime/web-server-demo/framework"
	"github.com/hktime/web-server-demo/framework/middlerware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	go func() {
		server.ListenAndServe()
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<- quit

	ctx, timeout := context.WithTimeout(context.Background(), 5 * time.Second)
	defer timeout()
	if err := server.Shutdown(ctx); err != nil{
		log.Fatal("server shutdown:", err)
	}
	select {
		case <-ctx.Done():
			log.Println("timeout 5s exceed")
	}
	log.Println("server exiting")
}
