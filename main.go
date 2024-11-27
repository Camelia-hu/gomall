package main

import (
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
)

func main() {
	conf.ViperInit()
	dao.MysqlInit()

}
