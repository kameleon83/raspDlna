package controllers

import (
	"text/template"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func Login() {
	sess := globalSessions.SessionStart()
	defer sess.SessionRelease()
	username := sess.Get("username")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		sess.Set("username", r.Form["username"])
	}
}
