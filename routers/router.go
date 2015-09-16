package routers

import (
	"fmt"
	"raspDlna/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/", beego.BeforeRouter, FilterLogin)
	beego.InsertFilter("/*", beego.BeforeRouter, FilterLogin)

	beego.Router("/login", &controllers.AuthController{}, "post,get:Login")
	beego.Router("/register", &controllers.AuthController{}, "post,get:Register")

	beego.Router("/", &controllers.ListController{})
	beego.Router("/list/:dir*", &controllers.ListController{})
	beego.Router("/vues/:files*.*", &controllers.VuesController{})
	beego.Router("/pas-vues/:files*.*", &controllers.PasVuesController{})
	beego.Router("/edit/:name*.*", &controllers.ChdirController{})
	beego.Router("/srt/:video*.*", &controllers.CmdController{}, "*:Srt")
	beego.Router("/delete/:f*.*", &controllers.DeleteController{}, "*:Delete")
	beego.Router("/rename/:old*.*", &controllers.CmdController{}, "*:Rename")

	//Erreurs
	beego.ErrorController(&controllers.ErrorController{})
}

var FilterLogin = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("name").(string)
	fmt.Println(ok)
	if !ok && ctx.Input.Uri() != "/login" && ctx.Input.Uri() != "/register" {
		ctx.Redirect(302, "/login")
	}
}
