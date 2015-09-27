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
		}
	}()
	c.Redirect("/list/"+path.Clean(path.Dir(lien)), 302)

}

func (c *CmdController) Mkdir() {
	lien := c.Ctx.Input.Param(":folder")
	folder := c.Ctx.Request.Form["mkdir"]
	a := path.Clean(lien + "/" + FormatString(strings.Join(folder, "")))
	d, _ := Emplacement(Root, a)
	if err := os.Mkdir(d, 0777); err != nil {
		check(err)
	}
	c.Redirect("/list/"+lien, 302)
}

func ListUser() []string {
	cmd := exec.Command("cat", "/etc/passwd")
	t, _ := cmd.Output()
	return strings.Split(string(t), "\n")
}

func ListGroup() []string {
	cmd := exec.Command("cat", "/etc/group")
	t, _ := cmd.Output()
	return strings.Split(string(t), "\n")
}

func (c *CmdController) Chown() {

	user := c.Ctx.Request.Form["user"]
	group := c.Ctx.Request.Form["group"]
	pass := c.Ctx.Request.Form["pass"]
	u := strings.Join(user, "")
	g := strings.Join(group, "")
	p := strings.Join(pass, "")
	go func() {
		cmd := exec.Command("/bin/sh", "-c", "sudo -S chown -R "+u+":"+g+" ~/Images/ <<< "+p)
		valeur, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erreur 1: ", err)
		} else {
			fmt.Println(string(valeur))
			fmt.Println("Chown effectué")
		}
	}()

	c.Redirect("/", 302)
}

func (c *CmdController) CmdPerso() {
	lien := c.Ctx.Input.Param("lien")
	// d, _ := Emplacement(Root, lien)
	c.Ctx.Request.ParseForm()
	args := c.Ctx.Request.Form["cmdperso"]
	a := strings.Join(args, "")
	go func() {
		cmd := exec.Command("/bin/sh", "-c", a)
		valeur, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erreur 1: ", err)
		} else {
			fmt.Println(string(valeur))
			fmt.Println("Commande effectué avec succès")
		}
	}()
	c.Redirect("/list/"+lien, 302)
}

func (c *CmdController) KeepOneAudio() {
	lien := c.Ctx.Input.Param(":video")
	d, _ := Emplacement(Root, lien)
	c.Ctx.Request.ParseForm()
	i := strings.Join(c.Ctx.Request.Form["number"], "")
	newName := strings.Join(c.Ctx.Request.Form["newName"], "")
	a, _ := Emplacement(Root, path.Dir(lien)+"/"+FormatString(newName))

	cmd := exec.Command("mkvmerge", "-o", a, "--atracks", i, d)
	go func() {
		t, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erreur : ", err)
		} else {
			fmt.Println(string(t))
			fmt.Println("Vidéo avec une seule piste audio")
		}
	}()
	c.Redirect("/list/"+path.Clean(path.Dir(lien)), 302)

}

func (c *CmdController) DtsToAc3() {
	flash := beego.NewFlash()
	lien := c.Ctx.Input.Param(":video")
	d, _ := Emplacement(Root, lien)
	c.Ctx.Request.ParseForm()
	newName := strings.Join(c.Ctx.Request.Form["dtstoac3"], "")
	a, _ := Emplacement(Root, path.Dir(lien)+"/"+FormatString(newName))
	cmd := exec.Command("ffmpeg", "-i", d, "-vcodec", "copy", "-scodec", "copy", "-acodec", "ac3", "-ac", "6", "-ab", "640k", a)
	go func() {
		t, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erreur : ", err)
		} else {
			fmt.Println(string(t))
			fmt.Println("Le son est maintenant en AC3")
		}
	}()
	flash.Notice("Encodage Du Dts en Ac3 en cours")
	flash.Store(&c.Controller)
	c.Redirect("/list/"+path.Clean(path.Dir(lien)), 302)
}

func SpaceDisk() []string {
	cmd := exec.Command("df", "-h", Root)
	b, _ := cmd.Output()
	return strings.Split(string(b), "\n")
}
func FormatString(file string) string {
	f := strings.ToLower(strings.Replace(file, " ", "-", -1))
	return f
}
