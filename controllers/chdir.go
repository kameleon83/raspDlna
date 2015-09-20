package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"raspDlna/models"
	"strings"

	"github.com/astaxie/beego"
)

type ChdirController struct {
	beego.Controller
}

func (c *ChdirController) Chdir() {
	edit := c.Ctx.Input.Param(":name")

	htmlMediainfo := mediaInfo(Root, edit)
	c.Data["name"] = path.Base(edit)
	c.Data["root"] = Root
	c.Data["htmlMediaInfo"] = htmlMediainfo
	c.Data["edit"] = edit
	t := models.ListFolder{}
	c.Data["listFolder"] = ReadJsonFolder(t, exepath, "listFolder")
	c.Data["back"] = path.Dir(edit)
	c.Layout = "index.tpl"
	c.TplNames = "chdir.tpl"
}

func (c *ChdirController) ChangeDir() {
	file := c.Ctx.Input.Param(":name")
	d, f := Emplacement(Root, file)
	finfo, err := os.Stat(d)

	go func() {
		c.Ctx.Request.ParseForm()
		Path := c.Ctx.Request.Form["newPath"]

		newPath := filepath.Clean(Root + strings.Join(Path, "") + "/" + f)

		if err != nil {
			check(err)
		} else {
			if !finfo.IsDir() {
				fmt.Println("c'est un fichier")
				Rename(d, newPath)
				fmt.Println("Le changement de répertoire est effectué")
			} else {
				Rename(d, newPath)
			}
		}
	}()
	c.Redirect("/list/"+path.Dir(file), 302)
}

func listFolder(root string) []string {
	fileList := []string{}
	len := len(root)
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			t := strings.TrimPrefix(path[len:], "/")

			fileList = append(fileList, t)
		}
		return nil
	})
	if err != nil {
		check(err)
	}
	return fileList
}

func mediaInfo(root, file string) string {
	d, _ := Emplacement(root, file)
	if finfo, err := os.Stat(d); err != nil {
		check(err)
	} else {
		if !finfo.IsDir() {

			cmd := exec.Command("mediainfo", "--Output=HTML", path.Clean(d))
			out, err := cmd.Output()
			if err != nil {
				fmt.Println("erreur 1 : ", path.Clean(root+file))
				check(err)
			}
			cmd.Wait()
			return string(out)
		}

	}
	return ""
}
