package models

type Configuration struct {
	Name     string `valid:"Required"`
	Password string `valid:"Required;Range(4,16)"`
	Root     string `valid:"Required"`
}
