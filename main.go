package main

import (
	_ "raspDlna/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func main() {
	globalSessions, _ = session.NewManager("memory",
		`{"cookieName":"gosessionid",
						"enableSetCookie,omitempty": true,
						"gclifetime":3600,
						"maxLifetime": 3600,
						"secure": false,
						"sessionIDHashFunc": "sha1",
						"sessionIDHashKey": "",
						"cookieLifeTime": 3600,
						 "providerConfig": ""}
						`)
	go globalSessions.GC()

	if beego.SessionOn != true {
		beego.SessionOn = true
		beego.SessionName = "raspDlna"
		beego.SessionProvider = "file"
		beego.SessionSavePath = "./tmp"
	}

	beego.Run()
}
