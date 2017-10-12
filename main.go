package main

import (
	"encoding/json"
	"log"
	_ "raspDlna/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func main() {
	beego.BConfig.AppName = "Server Files Movies"
	config := `{"cookieName" : "gosessionid",
						"enableSetCookie,omitempty": true,
						"gclifetime":3600,
						"maxLifetime": 3600,
						"secure": false,
						"sessionIDHashFunc": "sha1",
						"sessionIDHashKey": "",
						"cookieLifeTime": 3600,
						 "providerConfig": ""}
						`
	conf := new(session.ManagerConfig)
	if err := json.Unmarshal([]byte(config), conf); err != nil {
		log.Fatal("json decode error", err)
	}
	globalSessions, _ = session.NewManager("memory", conf)

	go globalSessions.GC()

	if beego.BConfig.WebConfig.Session.SessionOn != true {
		beego.BConfig.WebConfig.Session.SessionOn = true
		beego.BConfig.WebConfig.Session.SessionName = "raspDlna"
		beego.BConfig.WebConfig.Session.SessionProvider = "file"
		beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	}

	beego.Run()
}
