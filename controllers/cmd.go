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
	fmt.Println(lien)
	lien = html.UnescapeString(lien)
	fmt.Println(lien)
	d, _ := Emplacement(Root, lien)
	// // ffmpeg -i Movie.mkv -map 0:s:0 subs.srt
	dNew := strings.Replace(d, " ", `.`, -1)
	fileNotExt := strings.TrimSuffix(dNew, filepath.Ext(dNew))
	if result := Rename(d, dNew); !result {
		fmt.Println("Problème de renommage lors de l'extraction du srt")
	}

	cmd := exec.Command("ffmpeg", "-i", path.Clean(dNew), "-map", "0:s", path.Clean(fileNotExt)+".srt")
	go func() {
		err := cmd.Start()
		if err != nil {
			fmt.Println("Erreur : ", err)
		} else {
			fmt.Println("Le sous-titre est créé")
			fmt.Println(cmd)
		}
	}()
	c.Redirect("/list/"+path.Clean(path.Dir(lien)), 302)

}

func (c *CmdController) Rename() {
	oldFile := c.Ctx.Input.Param(":old")

	c.Ctx.Request.ParseForm()
	file := c.Ctx.Request.Form["rename"]
	newFile := filepath.Clean(strings.Replace(path.Dir(oldFile)+"/"+strings.Join(file, ""), " ", ".", -1))
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
	fmt.Println("It's false")
	return false
}
