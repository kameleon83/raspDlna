package routers

import (
	"raspDlna/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ListController{})
	beego.Router("/list/:dir*", &controllers.ListController{})
	beego.Router("/vues/:files*.*", &controllers.VuesController{})
	beego.Router("/pas-vues/:files*.*", &controllers.PasVuesController{})
	beego.Router("/edit/:name*.*", &controllers.ChdirController{})
	beego.Router("/srt/:video*.*", &controllers.CmdController{}, "*:Srt")
	beego.Router("/delete/:f*.*", &controllers.DeleteController{}, "*:Delete")
	beego.Router("/rename/:old*.*", &controllers.CmdController{}, "*:Rename")

}
