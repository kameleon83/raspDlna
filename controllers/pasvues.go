package controllers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type PasVuesController struct {
	beego.Controller
}

func (p *PasVuesController) Post() {

	file := p.Ctx.Input.Param(":files")
	d, f := Emplacement(Root, file)
	fileNotExt := strings.TrimSuffix(f, filepath.Ext(f))
	c := fileNotExt + ".srt"
	pathSrt := path.Dir(d) + "/" + fileNotExt + ".srt"
	if filepath.HasPrefix(f, ".") {
		finfo, err := os.Stat(d)
		if err != nil {
			check(err)
		} else {
			if !finfo.IsDir() {
				err := os.Rename(d, path.Dir(d)+"/"+strings.Replace(f, ".", "", 1))
				if err != nil {
					check(err)
				}
				if filepath.HasPrefix(c, ".") {
					fmt.Println(pathSrt, path.Dir(pathSrt)+"/"+strings.Replace(c, ".", "", 1))
					_, err = os.Stat(pathSrt)
					if err == nil {
						err := os.Rename(pathSrt, path.Dir(pathSrt)+"/"+strings.Replace(c, ".", "", 1))
						if err != nil {
							check(err)
						}
					}
				}
				p.Redirect("/list/"+path.Dir(file)+"/", 302)
			}
		}
	} else {
		p.Redirect("/list/"+path.Dir(file), 302)
	}
	p.Layout = "index.tpl"
	p.TplName = "list.tpl"
	p.Data["title"] = strings.Title("liste")
}
