package controllers

import (
	"fmt"
	"path"
	"raspDlna/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/kardianos/osext"
	. "golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller
}

func (l *AuthController) Login() {

	register := models.Configuration{}

	l.Ctx.Request.ParseForm()
	username := l.Ctx.Request.Form.Get("username")
	password := l.Ctx.Request.Form.Get("password")

	n, p, r, TorF := ReadJson(register, exepath)
	if !TorF {
		WriteJson(register, exepath, "config")
		l.Redirect("/register", 302)
	} else {
		if l.GetSession("name") != n {
			if l.Ctx.Input.Method() == "POST" {
				if err := l.ParseForm(&register); err != nil {
					fmt.Println("Pas de donnée parsé")
				} else {
					l.Ctx.Request.ParseForm()
					if username == n {
						if err := CompareHashAndPassword([]byte(p), []byte(password)); err != nil {
							fmt.Println("les mots de passes ne correspondent pas")
						} else {
							l.SetSession("name", username)
							l.SetSession("root", r)
							l.Redirect("/", 302)
						}

					} else {
						fmt.Println("les entrées ne sont pas correctes")
					}
				}
			}
		} else {
			l.Redirect("/", 302)
		}

	}

	l.Data["title"] = "Se connecter"
	l.Layout = "index.tpl"
	l.TplNames = "login.tpl"
}

func (l *AuthController) Register() {
	Root, _ := osext.ExecutableFolder()

	register := models.Configuration{}

	n, _, _, _ := ReadJson(register, Root)
	if l.GetSession("name") == n {
		l.Redirect("/", 302)
	} else {
		if l.Ctx.Input.Method() == "POST" {
			if err := l.ParseForm(&register); err != nil {
				fmt.Println("Pas de donnée parsé")
			} else {
				l.Ctx.Request.ParseForm()
				name := l.Ctx.Request.Form.Get("name")
				pass := l.Ctx.Request.Form.Get("password")
				confirmPass := l.Ctx.Request.Form.Get("confirmPassword")
				pathFolder := l.Ctx.Request.Form.Get("root")
				if pass == confirmPass {
					valid := validation.Validation{}
					if password, err := GenerateFromPassword([]byte(pass), DefaultCost); err == nil {
						register.Name = name
						register.Password = string(password)
						register.Root = path.Clean(pathFolder) + "/"
						if _, err := valid.Valid(&register); err != nil {
							fmt.Println("erreur de validation", err)
							fmt.Println(register)
						} else {
							WriteJson(register, Root, "config")
							l.Redirect("/", 302)
						}
					}
				}
			}
		}
	}

	l.Data["title"] = "Configurer l'application"
	l.Data["root"] = Root
	l.Layout = "index.tpl"
	l.TplNames = "register.tpl"
}
