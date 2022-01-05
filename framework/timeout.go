package framework

import (
	"context"
	"fmt"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler{
	return func(c *Context) error{
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d*time.Second)
		defer cancel()

		c.request.WithContext(durationCtx)
		go func() {
			defer func() {
				if p := recover(); p != nil{
					panicChan <- p
				}
			}()
			// 业务逻辑
			fun(c)

			finish <- struct{}{}
		}()
		select {
		case p := <-panicChan:
			fmt.Println(p)
			c.responseWriter.WriteHeader(500)
		case <- finish:
			fmt.Println("finish")
		case <- durationCtx.Done():
			c.responseWriter.Write([]byte("time out"))
		}
		return nil
	}
}

func Timeout(d time.Duration) ControllerHandler{
	return func(c *Context) error {
		return nil
	}
}