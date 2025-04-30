package main

import (
	_ "dev-assignment/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
