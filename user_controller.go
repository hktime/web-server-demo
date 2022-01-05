package main

import (
	"time"
	"web-demo/framework"
)

func UserLoginController(c *framework.Context) error {
	time.Sleep(1 * time.Second)
	c.Json(200, "ok, UserLoginController")
	return nil
}
