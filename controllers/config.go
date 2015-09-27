package controllers

import (
	"encoding/json"
	"io/ioutil"
	"raspDlna/models"

	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func ReadJson(config models.Configuration, exePath string) (string, string, string, bool) {
	file, err := ioutil.ReadFile(exePath + "/.config.json")
	if err != nil {
		return "", "", "", false
		panic("Impossible de transformer la configuration en JSON")
	}
	json.Unmarshal(file, &config)
	return config.Name, config.Password, config.Root, true
}

func WriteJson(config interface{}, exePath, nameFile string) {
	b, err := json.MarshalIndent(config, "", "  ")
	err = ioutil.WriteFile(exePath+"/."+nameFile+".json", b, 0755)
	if err != nil {
		panic("Impossible d'écrire dans le fichier de configuration")
	}
	if err != nil {
		panic("Problème avec le fichier de configuration")
	}
}

func ReadJsonFolder(config models.ListFolder, exePath, nameFile string) []string {
	file, err := ioutil.ReadFile(exePath + "/." + nameFile + ".json")
	if err != nil {
		panic("Impossible de transformer la configuration en JSON")
	}
	json.Unmarshal(file, &config)
	return config.Folder
}
