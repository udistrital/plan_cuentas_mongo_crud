package main

import (
	_ "github.com/udistrital/plan_cuentas_mongo_crud/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {

			ctx.Output.Header("Access-Control-Allow-Origin", "*")
			ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS")
			ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	})

	beego.Run()
}
