package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type ChdirController struct {
	beego.Controller
}

func (c *ChdirController) Get() {
	edit := c.Ctx.Input.Param(":name")
	lF := listFolder(Root, edit)
	htmlMediainfo := mediaInfo(Root, edit)
	c.Data["name"] = path.Base(edit)
	c.Data["root"] = Root
	c.Data["htmlMediaInfo"] = htmlMediainfo
	c.Data["edit"] = edit
	c.Data["listFolder"] = lF
	c.Data["back"] = path.Dir(edit)
	c.Layout = "index.tpl"
	c.TplNames = "chdir.tpl"
}

func (c *ChdirController) Post() {

	file := c.Ctx.Input.Param(":name")
	d, f := Emplacement(Root, file)

	c.Ctx.Request.ParseForm()
	Path := c.Ctx.Request.Form["newPath"]
	TextareaPath := c.Ctx.Request.Form["textarea-chemin"]
	newPath := filepath.Clean(Root + strings.Join(Path, "") + "/" + f)
	newPath2 := filepath.Clean(strings.Replace(strings.Join(TextareaPath, ""), " ", "-", -1))
	finfo, err := os.Stat(d)
	if err != nil {
		check(err)
	} else {
		if !finfo.IsDir() {
			go func() {
				if Path[0] == "" {
					a := filepath.Dir(newPath2)
					fmt.Println(a)
					chemin := strings.Replace(a, Root, "", -1)
					fmt.Println(chemin)
					bo := Rename(d, newPath2)
					if !bo {
						fmt.Println("erreur 2")
					}
					c.Redirect("/list/"+chemin, 302)
				} else {
					err := Rename(d, newPath)
					if !err {
						fmt.Println("erreur 3")
					}
					c.Redirect("/list/"+strings.Join(Path, ""), 302)
				}
				fmt.Println("Le changement de répertoire est effectué")
			}()
		}
	}

	c.Layout = "index.tpl"
	c.TplNames = "chdir.tpl"
}

func listFolder(root, file string) []string {
	fileList := []string{}
	len := len(root)
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			fileList = append(fileList, path[len:])
		}
		return nil
	})
	if err != nil {
		check(err)
	}
	return fileList
}

func mediaInfo(root, file string) string {
	cmd := exec.Command("mediainfo", "--Output=HTML", path.Clean(root+file))
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("erreur 1 : ", path.Clean(root+file))
		check(err)
	}
	cmd.Wait()
	return string(out)
}
