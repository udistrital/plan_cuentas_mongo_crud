package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	migrationmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/migrationManager"
	_ "github.com/udistrital/plan_cuentas_mongo_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
)

func main() {

	// beego.BConfig.RecoverFunc = responseformat.GlobalResponseHandler

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	if _, err := migrationmanager.RunMigrations(); err != nil {
		logs.Error("Migrations Error: ", err.Error())
	} else {
		logs.Info("Migration process success !")
	}

	//Prueba de auditoria
	// auditoria.InitMiddleware()

	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()
	beego.Run()
}
