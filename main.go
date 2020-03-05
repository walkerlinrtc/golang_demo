package main

import (
	_ "StitpProjectVersion1/routers"
	"github.com/astaxie/beego"
	"StitpProjectVersion1/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

