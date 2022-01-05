package main

import (
	"web-demo/framework"
	"web-demo/framework/middlerware"
)

func registerRouter(core *framework.Core) {
	core.Get("/user/login", middlerware.Test3(), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middlerware.Test3())
		subjectApi.Get("/:id", SubjectGetController)
	}
}
