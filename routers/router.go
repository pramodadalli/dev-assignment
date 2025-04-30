package routers

import (
	"dev-assignment/controllers"
	"dev-assignment/middlewares"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/register", &controllers.UserController{}, "post:Register")
	web.Router("/login", &controllers.UserController{}, "post:Login")

	web.Router("/upload", &controllers.FileController{}, "post:Upload")
	web.Router("/storage/remaining", &controllers.FileController{}, "get:RemainingStorage")
	web.Router("/files", &controllers.FileController{}, "get:ListFiles")

	web.InsertFilter("/upload", web.BeforeRouter, middlewares.JWTMiddleware)
	web.InsertFilter("/storage/remaining", web.BeforeRouter, middlewares.JWTMiddleware)
	web.InsertFilter("/files", web.BeforeRouter, middlewares.JWTMiddleware)
}
