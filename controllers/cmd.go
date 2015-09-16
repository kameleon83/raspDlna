package controllers

import (
	"fmt"
	"html"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type CmdController struct {
	beego.Controller
}

func (c *CmdController) Srt() {
	lien := c.Ctx.Input.Param(":video")
	lien = html.UnescapeString(lien)

	d, _ := Emplacement(Root, lien)
	fileNotExt := strings.TrimSuffix(d, filepath.Ext(d))
	// // ffmpeg -i Movie.mkv -map 0:s:0 subs.srt
	d = strings.Replace(d, " ", `\ `, -1)
	fileNotExt = strings.Replace(fileNotExt, " ", `\ `, -1)

	cmd := exec.Command("ffmpeg", "-i", path.Clean(d), "-map", "0:s:0", path.Clean(fileNotExt)+".srt")
	go func() {
		err := cmd.Start()
		if err != nil {
			fmt.Println("Erreur : ", err)
		}
		fmt.Println("Le sous-titre est créé")
	}()
	c.Redirect("/edit/"+lien, 302)

}

func (c *CmdController) Rename() {
	oldFile := c.Ctx.Input.Param(":old")

	c.Ctx.Request.ParseForm()
	file := c.Ctx.Request.Form["rename"]
	newFile := filepath.Clean(strings.Replace(path.Dir(oldFile)+"/"+strings.Join(file, ""), " ", "-", -1))
	dOld, _ := Emplacement(Root, oldFile)
	dNew, _ := Emplacement(Root, newFile)
	chemin := strings.Replace(path.Dir(dNew), Root, "", -1)
	c.Redirect("/list/"+chemin, 302)
	if result := Rename(dOld, dNew); result == true {
		fmt.Println("Le Dossier / fichier a bien été modifié")
	}
}

func Rename(oldFile, newFile string) bool {
	info, err := os.Stat(oldFile)
	if err != nil {
		return false
	}
	if info.IsDir() {
		if err := os.Mkdir(newFile, 0777); err != nil {
			check(err)
		} else {
			if err := os.Rename(oldFile, newFile); err != nil {
				check(err)
			} else {
				if v, err := Delete(oldFile); v != true && err != nil {
					check(err)
				} else {
					return true
				}
			}
		}
	} else {
		if err := os.Rename(oldFile, newFile); err != nil {
			check(err)
		} else {
			oldFileNotExt := strings.TrimSuffix(oldFile, filepath.Ext(oldFile))
			newFileNotExt := strings.TrimSuffix(newFile, filepath.Ext(newFile))
			oldFileWithExt := oldFileNotExt + ".srt"
			if _, err := os.Stat(oldFileWithExt); err != nil {
				check(err)
			} else {
				if err := os.Rename(oldFileWithExt, newFileNotExt+".srt"); err != nil {
					check(err)
				} else {
					return true
				}
			}

			return true
		}
	}

	return false
}
