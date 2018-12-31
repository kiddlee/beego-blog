package routers

import (
	"beegomyblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Include(&controllers.MainController{})
}
