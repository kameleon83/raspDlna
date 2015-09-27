package controllers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type RenameController struct {
	beego.Controller
}

func (c *RenameController) Rename() {
	oldFile := c.Ctx.Input.Param(":old")

	c.Ctx.Request.ParseForm()
	file := c.Ctx.Request.Form["rename"]
	newFile := filepath.Clean(path.Dir(oldFile) + "/" + FormatString(strings.Join(file, "")))
	dOld, _ := Emplacement(Root, oldFile)
	dNew, _ := Emplacement(Root, newFile)
	chemin := path.Clean(strings.Replace(path.Dir(dNew), Root, "", -1))
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
				return true
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
				fmt.Println("Il n'existe pas de sous titre")
				if f, err := os.Stat(path.Dir(oldFile)); err != nil {
					check(err)
				} else {
					if f.IsDir() {
						if v, err := IsEmpty(path.Dir(oldFile)); v == true && err == nil {
							if v, err := Delete(path.Dir(oldFile)); v != true && err != nil {
								check(err)
							} else {
								return true
							}
						}
					}
				}
				return true
			} else {
				if err := os.Rename(oldFileWithExt, newFileNotExt+".srt"); err != nil {
					check(err)
				} else {
					if f, err := os.Stat(path.Dir(oldFile)); err != nil {
						check(err)
					} else {
						if f.IsDir() {
							if v, err := IsEmpty(path.Dir(oldFile)); v == true && err == nil {
								if v, err := Delete(path.Dir(oldFile)); v != true && err != nil {
									check(err)
								} else {
									return true
								}
							}
						}
					}
					return true
				}
			}

			return true
		}
	}
	return false
}
