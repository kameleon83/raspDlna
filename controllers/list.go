package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"raspDlna/models"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/kardianos/osext"
)

type ListController struct {
	beego.Controller
}

type FileInfo struct {
	sync.Mutex
	Name     string
	Size     float64
	Mode     os.FileMode
	ModTime  time.Time
	IsDir    bool
	NameSize string
	Srt      int
}

var exepath, _ = osext.ExecutableFolder()
var _, _, Root, _ = ReadJson(models.Configuration{}, exepath)

func (c *ListController) Get() {
	fmt.Println(Root)
	flash := beego.ReadFromRequest(&c.Controller)
	if _, ok := flash.Data["notice"]; ok {
	}
	v := c.GetSession("raspDlna")
	if v == nil {
		c.SetSession("raspDlnaID", int(1))
		c.Data["num"] = 0

	} else {
		c.SetSession("raspDlnaID", v.(int)+1)
		c.Data["num"] = v.(int)
	}

	file := filepath.Clean(c.Ctx.Input.Param(":dir"))

	d, _ := Emplacement(Root, file)
	l := List(d)

	c.Data["title"] = strings.Title("liste")
	c.Data["href"] = "/" + file
	c.Data["back"] = path.Dir(file)
	c.Data["chemin"] = strings.Split(file, "/")

	// c.Data["breadcrumb"] = b
	c.Data["dirname"] = l
	c.Layout = "index.tpl"
	c.TplNames = "list.tpl"
}

func Emplacement(dir, file string) (dire, fileExt string) {
	fileExt = filepath.Clean(path.Base(file))
	dire = filepath.Clean(path.Dir(dir + file))
	if dire != "" {
		dire = dir + "/" + file
	}
	return filepath.Clean(dire), fileExt
}

func List(dir string) []FileInfo {
	list := []FileInfo{}
	fi := FileInfo{}
	if info, err := os.Stat(dir); err == nil && info.IsDir() {
		f, _ := ioutil.ReadDir(dir)
		for _, entry := range f {
			namePresent := ""
			fileNotExt := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
			taille, nomTaille := Taille(float64(entry.Size()))
			if a, _ := SrtOrNot(entry.Name()); a != true {
				for _, entry2 := range f {
					if a, b := SrtOrNot(entry2.Name()); a == true {
						if fileNotExt == b {
							namePresent = entry.Name()
							fi = FileInfo{
								Name:     entry.Name(),
								Size:     float64(int(taille*100)) / 100,
								Mode:     entry.Mode(),
								ModTime:  entry.ModTime(),
								IsDir:    entry.IsDir(),
								NameSize: nomTaille,
								Srt:      1,
							}
							fi.Lock()
							list = append(list, fi)
							fi.Unlock()
						}
					}
				}
				if namePresent != entry.Name() {

					fi = FileInfo{
						Name:     entry.Name(),
						Size:     float64(int(taille*100)) / 100,
						Mode:     entry.Mode(),
						ModTime:  entry.ModTime(),
						IsDir:    entry.IsDir(),
						NameSize: nomTaille,
						Srt:      0,
					}
					fi.Lock()
					list = append(list, fi)
					fi.Unlock()
				}

			}
		}
	}
	return list
}

func SrtOrNot(file string) (bool, string) {
	ext := path.Ext(file)
	fileWithExt := strings.TrimSuffix(file, filepath.Ext(file))
	if ext == ".srt" {
		return true, fileWithExt
	}
	return false, ""
}

func Taille(i float64) (float64, string) {
	taille := 0.00
	nomTaille := ""
	if i > 1073741824 {
		taille = float64(i) / 1073741824
		nomTaille = "Go"
	} else if i > 1048576 {
		taille = float64(i) / 1048576
		nomTaille = "Mo"
	} else if i > 1024 {
		taille = float64(i) / 1024
		nomTaille = "Ko"
	} else {
		taille = float64(i)
		nomTaille = "o"
	}
	return taille, nomTaille
}
