package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"raspDlna/models"

	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func ReadJson(config models.Configuration, exePath string) {
	file, err := ioutil.ReadFile(exePath + "/.config.json")
	if err != nil {
		panic("Impossible de transformer la configuration en JSON")
	}
	json.Unmarshal(file, &config)
	fmt.Printf("Port : %s\n", config.Port)
	fmt.Printf("Folders : %s\n", config.Folders)
}

func WriteJson(config models.Configuration, exePath string) {

	b, err := json.MarshalIndent(config, "", "  ")
	err = ioutil.WriteFile(exePath+"/.config.json", b, 0644)
	if err != nil {
		panic("Impossible d'écrire dans le fichier de configuration")
	}
	if err != nil {
		panic("Problème avec le fichier de configuration")
	}
}
