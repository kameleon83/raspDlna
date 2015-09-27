package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"raspDlna/models"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/astaxie/beego"
	"github.com/kardianos/osext"
)

type ListController struct {
	beego.Controller
}

var exepath, _ = osext.ExecutableFolder()
var _, _, Root, _ = ReadJson(models.Configuration{}, exepath)

func (c *ListController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)

	if _, err := os.Stat(exepath + "/.config.json"); err != nil {
		fmt.Println("erreur Il n'y a pas de config.json")
		c.DestroySession()
	}

	go func() {
		t := models.ListFolder{}
		t.Folder = listFolder(Root)
		WriteJson(t, exepath, "listFolder")
	}()

	if _, ok := flash.Data["notice"]; ok {
	}
	if _, ok := flash.Data["error"]; ok {
	}
	if _, ok := flash.Data["success"]; ok {
	}
	if _, ok := flash.Data["warning"]; ok {
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

	user, group := MapUidGid()
	c.Data["user"] = user
	c.Data["group"] = group

	sd := SpaceDisk()
	reg, _ := regexp.Compile("[^A-Za-z0-9/-]+")
	sd1 := strings.Split(reg.ReplaceAllString(sd[1], " "), " ")
	c.Data["spacedisk"] = sd1

	// c.Data["breadcrumb"] = b
	c.Data["dirname"] = l
	c.Layout = "index.tpl"
	c.TplNames = "list.tpl"
}

func MapUidGid() (map[string]int, map[string]int) {
	u := make(map[string]int)
	for _, v := range ListUser() {
		a := strings.Split(v, ":")
		if a[0] != "" {
			s, _ := strconv.Atoi(a[2])
			u[a[0]] = s
		}
	}
	g := make(map[string]int)
	for _, v := range ListGroup() {
		a := strings.Split(v, ":")
		if a[0] != "" {
			s, _ := strconv.Atoi(a[2])
			g[a[0]] = s
		}
	}
	return u, g
}

func Emplacement(dir, file string) (dire, fileExt string) {
	fileExt = filepath.Clean(path.Base(file))
	dire = filepath.Clean(path.Dir(dir + file))
	if dire != "" {
		dire = dir + "/" + file
	}
	return filepath.Clean(dire), fileExt
}

func List(dir string) []models.FileInfo {
	list := []models.FileInfo{}
	fi := models.FileInfo{}
	if info, err := os.Stat(dir); err == nil && info.IsDir() {
		f, _ := ioutil.ReadDir(dir)
		for _, entry := range f {
			namePresent := ""
			fileNotExt := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
			taille, nomTaille := Taille(float64(entry.Size()))
			if a, _ := SrtOrNot(entry.Name()); a != true {
				for _, entry2 := range f {
					tailleSrt, SrtNametaillle := Taille(float64(entry2.Size()))
					if a, b := SrtOrNot(entry2.Name()); a == true {
						if fileNotExt == b {
							namePresent = entry.Name()
							fi = models.FileInfo{
								Name:          entry.Name(),
								NameExt:       filepath.Ext(entry.Name()),
								Size:          float64(int(taille*100)) / 100,
								Mode:          entry.Mode(),
								ModTime:       entry.ModTime(),
								IsDir:         entry.IsDir(),
								NameSize:      nomTaille,
								Srt:           1,
								SizeSrt:       float64(int(tailleSrt*100)) / 100,
								NameTailleSrt: SrtNametaillle,
								GetUid:        entry.Sys().(*syscall.Stat_t).Uid,
								GetGid:        entry.Sys().(*syscall.Stat_t).Gid,
							}
							fi.Lock()
							list = append(list, fi)
							fi.Unlock()
						}
					}
				}
				if namePresent != entry.Name() {

					fi = models.FileInfo{
						Name:     entry.Name(),
						NameExt:  filepath.Ext(entry.Name()),
						Size:     float64(int(taille*100)) / 100,
						Mode:     entry.Mode(),
						ModTime:  entry.ModTime(),
						IsDir:    entry.IsDir(),
						NameSize: nomTaille,
						Srt:      0,
						GetUid:   entry.Sys().(*syscall.Stat_t).Uid,
						GetGid:   entry.Sys().(*syscall.Stat_t).Gid,
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
