package controllers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type VuesController struct {
	beego.Controller
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (v *VuesController) Post() {

	file := v.Ctx.Input.Param(":files")
	d, f := Emplacement(beego.AppConfig.String("homeDirectory"), file)
	fileNotExt := strings.TrimSuffix(f, filepath.Ext(f))
	c := fileNotExt + ".srt"
	pathSrt := path.Dir(d) + "/" + fileNotExt + ".srt"

	if !filepath.HasPrefix(f, ".") {
		finfo, err := os.Stat(d)
		if err != nil {
			check(err)
		} else {
			if !finfo.IsDir() {
				err := os.Rename(d, path.Dir(d)+"/."+f)
				if err != nil {
					check(err)
				}
				if !filepath.HasPrefix(c, ".") {
					_, err = os.Stat(pathSrt)
					if err == nil {
						err := os.Rename(pathSrt, path.Dir(pathSrt)+"/."+c)
						if err != nil {
							check(err)
						}
					}
				}
				v.Redirect("/list/"+path.Dir(file), 302)
			}
		}
	} else {
		fmt.Println(" le fichier a déjà été modifié !")
		v.Redirect("/list/"+path.Dir(file), 302)
	}
}
