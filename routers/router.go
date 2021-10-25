package routers

import (
	"leonet-gateway/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/test", &controllers.MainController{})
}
