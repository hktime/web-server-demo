package main

import "web-demo/framework"

func SubjectGetController(c *framework.Context)error{
	c.Json(200,"get subject")
	return nil
}
