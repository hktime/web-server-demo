package main

import (
	"time"
	"web-demo/framework"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	time.Sleep(10 * time.Second)
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
