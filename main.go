package main

import (
	_ "leonet-gateway/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

