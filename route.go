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
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middlerware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
