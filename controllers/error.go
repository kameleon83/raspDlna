package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	flash := beego.NewFlash()
	flash.Error("Erreur : la page demandée n'est pas accessible")
	flash.Store(&c.Controller)
	c.Data["content"] = "page not found"
	c.Redirect("/", 302)
}

// func (c *ErrorController) Error501() {
// 	c.Data["content"] = "server error"
// 	c.TplNames = "501.tpl"
// }

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplNames = "error/dbError.tpl"
}
