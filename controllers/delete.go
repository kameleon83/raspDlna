package controllers

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

// DeleteController : Structure
type DeleteController struct {
	beego.Controller
}

// Delete : Fonction qui est appelé
func (d *DeleteController) Delete() {
	file := d.Ctx.Input.Param(":f")
	f := path.Dir(file)
	p, _ := Emplacement(Root, file)
	fmt.Println(p)
	d.Redirect("/list/"+f, 302)
	r, err := Delete(p)
	if r != true || err != nil {
		check(err)
	}
}

// Delete : FOnction de suppresion des fichiers et / ou des dossiers
func Delete(pathFile string) (bool, error) {
	finfo, err := os.Stat(pathFile)
	if err != nil {
		fmt.Println("test dans le delete")
		check(err)
	} else {
		if finfo.IsDir() {
			if v, err := IsEmpty(pathFile); v == true && err == nil {
				err := os.Remove(pathFile)
				if err != nil {
					check(err)
				} else {
					return true, nil
				}
			} else {
				fmt.Println("Attention le Dossier n'est pas vide. Par conséquent la suppresion n'est pas possible")
				fmt.Println("Cette option sera arrivera bientôt")
				// os.RemoveAll(pathFile)
			}
		} else {
			err := os.Remove(pathFile)
			if err != nil {
				check(err)
			} else {
				pathFileNotExt := strings.TrimSuffix(pathFile, filepath.Ext(pathFile))
				pathFileWithExt := pathFileNotExt + ".srt"
				if _, err := os.Stat(pathFileWithExt); err != nil {
					check(err)
				} else {
					if err := os.Remove(pathFileWithExt); err != nil {
						check(err)
					} else {
						return true, nil
					}
				}
				return true, nil
			}

		}
	}
	return false, err
}

// IsEmpty : Vérification si le dossier est vide ou pas
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
