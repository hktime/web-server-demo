package main

import "web-demo/framework"

func registerRouter(core *framework.Core) {
	core.Get("/foo", FooControllerHandler)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Get("/:id", SubjectGetController)
	}
}
